package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/buy_event/config"
	"github.com/buy_event/storage/repo"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"strings"
)

// purchaseCmd represents the purchase command
var purchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase command for creating, updating, deleting, getting purchases",
}

var addPurchaseCmd = &cobra.Command{
	Use: "add",
	Short:"create a new purchase",
	RunE: func(cmd *cobra.Command, args []string) error {
		userID, err := cmd.Flags().GetString("user-id")

		if err != nil {
			return err
		}

		if userID == "" {
			return errors.New("user-id not found")
		}

		userUID, err := uuid.Parse(userID)

		if err != nil {
			return err
		}

		price, err := cmd.Flags().GetFloat64("price")

		if err != nil {
			return err
		}

		if price < 0 {
			return errors.New("price should be non negative number")
		}

		goods, _ := cmd.Flags().GetStringArray("good")

		if len(goods) == 0 {
			if len(args) == 0 {
				return errors.New("goods not found")
			}
			goods = args
		}

		err = purchaseService.Create(context.Background(), &repo.Purchase{
			Goods: strings.Join(goods, ","),
			UserID: userUID,
			Price:price,
		})

		if err != nil {
			return err
		}

		cmd.Println("purchase created successfully")

		messageType, err := cmd.Flags().GetString("message")

		if err != nil {
			return err
		}

		if !validateMessageType(messageType) {
			_ = logService.Create(context.Background(), &repo.Log{
				Error: fmt.Sprintf("message sending type not found. type: %s", messageType),
			})
			return errors.New("message type not found")
		}

		user, err := userService.Get(context.Background(), userUID)

		if err != nil {
			return err
		}
		if messageType == config.MESSAGE_TYPE_SMS {
			err = sendSMS(user.PhoneNumber, fmt.Sprintf("You bought %s goods and total price is %.2f", strings.Join(goods, ","), price))

			if err != nil {
				_ = logService.Create(context.Background(), &repo.Log{
					Error: err.Error(),
				})

				return errors.New("notification did not send")
			}
			cmd.Println("notification successfully sent")
		} else {
			err = sendEmail(user.Email, fmt.Sprintf("You bought %s goods and total price is %.2f", strings.Join(goods, ","), price))

			if err != nil {
				_ = logService.Create(context.Background(), &repo.Log{
					Error: err.Error(),
				})

				return errors.New("notification did not send")
			}
			cmd.Println("notification successfully sent")
		}
		return nil
	},
}

var getPurchaseCmd = &cobra.Command{
	Use: "get",
	Short: "Get a purchase",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")

		if err != nil {
			return err
		}

		if id == "" {
			purchases, err := purchaseService.GetAll(context.Background())

			if err != nil {
				return err
			}

			for _, purchase := range purchases {
				cmd.Println(`----------**********----------`)
				cmd.Println("ID: ", purchase.ID)
				cmd.Println("Goods: ", purchase.Goods)
				cmd.Println("Price: ", purchase.Price)
				cmd.Println("User ID: ", purchase.UserID)
			}
		} else {
			uID, err := uuid.Parse(id)

			if err != nil {
				return err
			}

			purchase, err := purchaseService.Get(context.Background(), uID)

			if err != nil {
				return err
			}

			cmd.Println("ID: ", purchase.ID)
			cmd.Println("Goods: ", purchase.Goods)
			cmd.Println("Price: ", purchase.Price)
			cmd.Println("User ID: ", purchase.UserID)
		}

		return nil
	},
}

var deletePurchaseCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete purchase",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")

		if err != nil {
			return err
		}

		if id == "" {
			if len(args) == 0{
				return errors.New("purchase-id not found")
			}

			id =args[0]
		}
		uid, err := uuid.Parse(id)

		if err != nil {
			return err
		}

		err = purchaseService.Delete(context.Background(), uid)

		if err != nil {
			return err
		}

		cmd.Println("purchase successfully deleted")
		return nil
	},
}


func init() {
	addPurchaseCmd.Flags().StringP("user-id", "u", "", "Enter user ID")
	addPurchaseCmd.Flags().Float64P("price", "p", -1, "Enter price")
	addPurchaseCmd.Flags().StringArrayP("good", "g", []string{}, "Enter good")
	addPurchaseCmd.Flags().StringP("message", "m", "email", "Message sending type")

	getPurchaseCmd.Flags().StringP("id", "i", "", "Enter purchase ID")

	deletePurchaseCmd.Flags().StringP("id", "i", "", "Enter purchase ID")

	purchaseCmd.AddCommand(addPurchaseCmd)
	purchaseCmd.AddCommand(getPurchaseCmd)
	purchaseCmd.AddCommand(getPurchaseCmd)
	purchaseCmd.AddCommand(deletePurchaseCmd)

	rootCmd.AddCommand(purchaseCmd)
}

func validateMessageType(mType string) bool {
	for _, val := range config.MessageTypes {
		if val == mType {
			return true
		}
	}
	return false
}
