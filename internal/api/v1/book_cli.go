package v1

import (
	"context"
	"fmt"
	"github.com/GGmaz/BookManager/internal/service"
	"github.com/GGmaz/BookManager/pkg/requests"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func registerBookCLI(ctx context.Context) {
	bookService, err := service.NewBook()
	if err != nil {
		log.WithFields(log.Fields{"service_name": "book-service", "method_name": "registerBookCLI"}).Error("Error creating book service: ", err)
		panic("Error creating book service.")
	}
	log.WithFields(log.Fields{"service_name": "book-service", "method_name": "registerBookCLI"}).Info("Successfully created book handler.")

	h := &BookHandler{
		bookService: bookService,
		ctx:         ctx,
	}

	app := &cli.App{
		Name:  "bookmanager-cli",
		Usage: "CLI client for BookManager service",
	}

	createCmd, getAllCmd, updateCmd, deleteCmd := createCliComands(h)

	app.Commands = append(app.Commands, createCmd, getAllCmd, updateCmd, deleteCmd)

	err = app.Run(os.Args)
	if err != nil {
		return
	}
}

func createCliComands(h *BookHandler) (*cli.Command, *cli.Command, *cli.Command, *cli.Command) {
	createCmd := &cli.Command{
		Name:  "create",
		Usage: "Create a new book",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Usage:    "Title of the book",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "author",
				Usage: "Author of the book",
			},
			&cli.StringFlag{
				Name:  "publishedDate",
				Usage: "Published date of the book",
			},
			&cli.StringFlag{
				Name:  "edition",
				Usage: "Edition of the book",
			},
			&cli.StringFlag{
				Name:  "description",
				Usage: "Description of the book",
			},
			&cli.StringFlag{
				Name:  "genre",
				Usage: "Genre of the book",
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("title")
			author := c.String("author")
			publishedDate := c.String("publishedDate")
			edition := c.String("edition")
			description := c.String("description")
			genre := c.String("genre")

			return h.CreateCLI(title, author, publishedDate, edition, description, genre)
		},
	}

	getAllCmd := &cli.Command{
		Name:  "get-all",
		Usage: "Get all books",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "page",
				Usage: "Page number",
			},
			&cli.IntFlag{
				Name:  "pageSize",
				Usage: "Page size",
			},
			&cli.StringFlag{
				Name:  "author",
				Usage: "Author filter",
			},
			&cli.StringFlag{
				Name:  "genre",
				Usage: "Genre filter",
			},
			&cli.StringFlag{
				Name:  "startDate",
				Usage: "Start date filter",
			},
			&cli.StringFlag{
				Name:  "endDate",
				Usage: "End date filter",
			},
		},
		Action: func(c *cli.Context) error {
			page := c.Int("page")
			pageSize := c.Int("pageSize")
			author := c.String("author")
			genre := c.String("genre")
			startDate := c.String("startDate")
			endDate := c.String("endDate")

			return h.GetAllCLI(page, pageSize, author, genre, startDate, endDate)
		},
	}

	updateCmd := &cli.Command{
		Name:  "update",
		Usage: "Update a book",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "id",
				Usage:    "ID of the book to update",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "title",
				Usage: "New title of the book",
			},
			&cli.StringFlag{
				Name:  "author",
				Usage: "New author of the book",
			},
			&cli.StringFlag{
				Name:  "publishedDate",
				Usage: "New published date of the book",
			},
			&cli.StringFlag{
				Name:  "edition",
				Usage: "New edition of the book",
			},
			&cli.StringFlag{
				Name:  "description",
				Usage: "New description of the book",
			},
			&cli.StringFlag{
				Name:  "genre",
				Usage: "New genre of the book",
			},
		},
		Action: func(c *cli.Context) error {
			id := c.Int64("id")
			title := c.String("title")
			author := c.String("author")
			publishedDate := c.String("publishedDate")
			edition := c.String("edition")
			description := c.String("description")
			genre := c.String("genre")

			return h.UpdateCLI(id, title, author, publishedDate, edition, description, genre)
		},
	}

	deleteCmd := &cli.Command{
		Name:  "delete",
		Usage: "Delete a book",
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:     "id",
				Usage:    "ID of the book to delete",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			id := c.Int64("id")

			return h.DeleteCLI(id)
		},
	}
	return createCmd, getAllCmd, updateCmd, deleteCmd
}

// CreateCLI handles the creation of a new book from the CLI.
func (handler *BookHandler) CreateCLI(title, author, publishedDate, edition, description, genre string) error {
	date, _ := time.Parse("2006-01-02T15:04:05Z", publishedDate)

	bookRequest := requests.CreateBookRequest{
		Title:         title,
		Author:        author,
		Edition:       edition,
		Description:   description,
		Genre:         genre,
		PublishedDate: date,
	}

	id := handler.bookService.Create(bookRequest)

	if id == 0 {
		fmt.Println("Error while creating a book.")
		os.Exit(1)
	} else {
		fmt.Println("Book successfully created with ID:", id)
	}

	os.Exit(0)
	return nil
}

// GetAllCLI handles retrieving all books from the CLI.
func (handler *BookHandler) GetAllCLI(page, pageSize int, author, genre, startDate, endDate string) error {
	books, total := handler.bookService.GetAll(page, pageSize, author, genre, startDate, endDate)

	fmt.Println("Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}

	fmt.Println("Total:", total)
	fmt.Println("Page:", page)

	os.Exit(0)
	return nil
}

// UpdateCLI handles updating a book from the CLI.
func (handler *BookHandler) UpdateCLI(id int64, title, author, publishedDate, edition, description, genre string) error {
	date, _ := time.Parse("2006-01-02T15:04:05Z", publishedDate)

	bookRequest := requests.UpdateBookRequest{
		Id:            id,
		Title:         title,
		Author:        author,
		PublishedDate: date,
		Edition:       edition,
		Description:   description,
		Genre:         genre,
	}

	errMessage := handler.bookService.Update(bookRequest.Id, bookRequest)

	if errMessage != "" {
		fmt.Println("Error:", errMessage)
		os.Exit(1)
	} else {
		fmt.Println("Book successfully updated.")
	}

	os.Exit(0)
	return nil
}

// DeleteCLI handles deleting a book from the CLI.
func (handler *BookHandler) DeleteCLI(id int64) error {
	errMessage := handler.bookService.Delete(id)

	if errMessage != "" {
		fmt.Println("Error:", errMessage)
		os.Exit(1)
	} else {
		fmt.Println("Book successfully deleted.")
	}

	os.Exit(0)
	return nil
}
