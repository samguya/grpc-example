package main
import (
    pb "github.com/samguya/grpc-example/protos/v1/user"
    "context"
    "google.golang.org/grpc"
    "log"
    "time"
    "fmt"
)   
    
func main() {
    conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewUserClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
        
    r, err := c.GetUser(ctx, &pb.GetUserRequest{UserId: "1"})
    //r, err := c.ListUsers(ctx, &pb.ListUsersRequest{})
    if err != nil {
        log.Fatalf("could not request: %v", err)
    }
    fmt.Println(r)
    log.Printf("Config: %v", r)
}
