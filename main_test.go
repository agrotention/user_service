package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/agrotention/user_proto"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Fungsi untuk mendapatkan klien gRPC yang valid
func getClient() user_proto.UserServiceClient {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	client := user_proto.NewUserServiceClient(conn)
	return client
}

// Fungsi untuk menguji koneksi ke server gRPC
func TestClient(t *testing.T) {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err.Error()) // Gagal jika tidak bisa menghubungkan
	}
	// Pastikan klien bisa dibuat tanpa error
	client := user_proto.NewUserServiceClient(conn)
	if client == nil {
		t.Fatal("Failed to create UserService client")
	}
}

// Fungsi untuk menguji registrasi pengguna yang sukses
func TestSuccessRegister(t *testing.T) {
	// Membuat request registrasi dengan data yang valid
	request := user_proto.InUserRegister{
		Username: fmt.Sprintf("testuser"),
		Password: "password123",
		FullName: "Test User",
	}

	// Mendapatkan klien gRPC
	client := getClient()
	ctx := context.Background()

	// Mengirim permintaan ke server untuk registrasi
	res, err := client.UserRegister(ctx, &request)
	if err != nil {
		t.Fatal("Failed to register user:", err.Error())
	}
	t.Logf("User %s created successfully!", res.GetId())

}
