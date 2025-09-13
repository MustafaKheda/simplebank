package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/MustafaKheda/simplebank/api"
	db "github.com/MustafaKheda/simplebank/db/sqlc"
	_ "github.com/MustafaKheda/simplebank/doc/statik"
	"github.com/MustafaKheda/simplebank/gapi"
	"github.com/MustafaKheda/simplebank/pb"
	"github.com/MustafaKheda/simplebank/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Connot Load Config File")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot Create Connection")
	}
	// run db migration
	runDbMigration(config.MigrationURL, config.DBSource)
	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}
func runDbMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create migration instance: ", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migation up: ", err)
	}
	log.Println("db migrated successfully")
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Create Server: ", err)

	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot Create Listener: ", err)
	}
	log.Printf("Start gRPC Server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot Start gRPC Server: ", err)
	}
}
func runGatewayServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Create Server: ", err)

	}
	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOptions)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Cannot Register Handler Server: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	// fs := http.FileServer(http.Dir("./doc/swagger/"))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal("Cannot Create statik file system: ", err)
	}
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(statikFS)))
	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot Create Listener: ", err)
	}
	log.Printf("Start Http Gateway Server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Cannot Start HTTP Gateway Server: ", err)
	}
}
func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Create Server: ", err)

	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("Cannot Start server")
	}
}
