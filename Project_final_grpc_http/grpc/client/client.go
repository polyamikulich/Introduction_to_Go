package main

import (
	"Go_Project/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new_name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := deleteCl(cmd); err != nil {
			return fmt.Errorf("delete account failed %w", err)
		}

		return nil
	case "patch":
		if err := patch(cmd); err != nil {
			return fmt.Errorf("patch account failed %w", err)
		}

		return nil
	case "change":
		if err := change(cmd); err != nil {
			return fmt.Errorf("change account failed %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func create(cmd Command) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.CreateAccount(ctx, &proto.CreateAccountReq{Username: cmd.Name, Amount: int32(cmd.Amount)})
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return fmt.Errorf("error: username %s already exist\n", cmd.Name)
		}
		return fmt.Errorf("error %w:", err)
	}

	return nil
}

func get(cmd Command) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetAccount(ctx, &proto.GetAccountReq{Username: cmd.Name})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return fmt.Errorf("error: username %s does not exist\n", cmd.Name)
		}

		return fmt.Errorf("error: %w", err)
	}

	fmt.Println(resp)

	return nil
}

func deleteCl(cmd Command) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.DeleteAccount(ctx, &proto.DeleteAccountReq{Username: cmd.Name})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return fmt.Errorf("error: username %s does not exist\n", cmd.Name)
		}

		return fmt.Errorf("error: %w", err)
	}

	fmt.Printf("Deleted: %s \n", cmd.Name)

	return nil
}

func patch(cmd Command) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.PatchAccount(ctx, &proto.PatchAccountReq{Username: cmd.Name, Amount: int32(cmd.Amount)})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return fmt.Errorf("account with name %s not founded", cmd.Name)
		default:
			return fmt.Errorf("error: %w", err)
		}
	}

	fmt.Printf("Patched: %s \n", cmd.Name)

	return nil
}

func change(cmd Command) error {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.ChangeAccount(ctx, &proto.ChangeAccountReq{LastName: cmd.Name, NewName: cmd.NewName})
	if err != nil {
		switch status.Code(err) {
		case codes.AlreadyExists:
			return fmt.Errorf("error: username %s already exist\n", cmd.NewName)
		case codes.NotFound:
			return fmt.Errorf("error: username %s does not exist\n", cmd.Name)
		default:
			return fmt.Errorf("error: %w", err)
		}
	}

	fmt.Printf("Changed: %s->%s \n", cmd.Name, cmd.NewName)

	return nil
}
