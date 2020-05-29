package cmd

import (
	"context"
	"errors"
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// addUserCmd represents the addUser command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User command for adding, getting, deleting users",
}

var addUserCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new user",
	RunE: func(cmd *cobra.Command, args []string) error {
		email, err := cmd.Flags().GetString("email")

		if err != nil {
			return err
		}

		if email == "" {
			return errors.New("user email not found")
		}

		phone, err := cmd.Flags().GetString("phone")

		if phone == "" {
			return errors.New("user phone not found")
		}

		if err != nil {
			return err
		}

		err = userService.Create(context.Background(), &repo.User{
			Email:email,
			PhoneNumber:phone,
		})

		if err != nil {
			return errors.New("error while creating a new user")
		}

		cmd.Println("user successfully created")
		return nil
	},
}

var getUserCmd = &cobra.Command{
	Use: "get",
	Short: "Get a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")

		if err != nil {
			return err
		}

		if id == "" {
			users, err := userService.GetAll(context.Background())

			if err != nil {
				return err
			}

			for _, user := range users {
				cmd.Println(`----------**********----------`)
				cmd.Println("ID: ", user.ID)
				cmd.Println("Email: ", user.Email)
				cmd.Println("Phone: ", user.PhoneNumber)
			}
		} else {
			uID, err := uuid.Parse(id)

			if err != nil {
				return err
			}

			user, err := userService.Get(context.Background(), uID)

			if err != nil {
				return err
			}

			cmd.Println("ID: ", user.ID)
			cmd.Println("Email: ", user.Email)
			cmd.Println("Phone: ", user.PhoneNumber)
		}

		return nil
	},
}

var deleteUserCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete user",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")

		if err != nil {
			return err
		}

		if id == "" {
			if len(args) == 0{
				return errors.New("user-id not found")
			}

			id =args[0]
		}
		uid, err := uuid.Parse(id)

		if err != nil {
			return err
		}

		err = userService.Delete(context.Background(), uid)

		if err != nil {
			return err
		}

		cmd.Println("user successfully deleted")
		return nil
	},
}

func init() {
	addUserCmd.Flags().StringP("email", "e", "", "Enter user email")
	addUserCmd.Flags().StringP("phone", "p", "", "Enter user phone")

	getUserCmd.Flags().StringP("id", "i", "", "Enter user ID")

	deleteUserCmd.Flags().StringP("id", "i", "", "Enter user id")

	userCmd.AddCommand(addUserCmd)
	userCmd.AddCommand(getUserCmd)
	userCmd.AddCommand(deleteUserCmd)
	rootCmd.AddCommand(userCmd)
}
