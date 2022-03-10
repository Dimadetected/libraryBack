package handler

import (
	"fmt"
	"github.com/Dimadetected/libraryBack/internal/models"
	"github.com/Dimadetected/libraryBack/pkg/respfmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) BooksGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	books, err := h.uc.GetBooks(limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, books)
}
func (h *Handler) BookGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	book, err := h.uc.GetBook(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, book)
}
func (h *Handler) BooksDelete(c *gin.Context) {
	setCors(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err := h.uc.DeleteBook(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	respfmt.OK(c, id)
}
func (h *Handler) BooksUpdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	id, err = h.uc.UpdateBook(id, book)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) BooksFilesUpdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	formFile, err := c.FormFile("file_upload")
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	file, err := formFile.Open()
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
		return
	}

	if err = h.uc.UpdateBookFile(id, file); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) BooksStore(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := h.uc.StoreBook(book)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) BooksTagsAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.BookTags

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksTagsAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) BooksTagsDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.BookTags

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksTagsDelete(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) FavoriteBooksAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.FavoriteBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}
	fmt.Printf("1 %+v\n", bookTags)
	if err := h.uc.FavoriteBooksAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) FavoriteBooksGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	favorites := make([]models.FavoriteBook, 1)
	favorites, err = h.uc.FavoriteBooksGet(id)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}
	fmt.Printf("%+v\n", favorites)
	respfmt.OK(c, favorites)
}
func (h *Handler) FavoriteBooksDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.FavoriteBooksDelete(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) ProcessingBooksAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var bookTags models.ProcessingBook

	if err := c.BindJSON(&bookTags); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.ProcessingBooksAdd(bookTags); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) ProcessingBooksGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
	}

	pb := make([]models.ProcessingBook, 0)
	pb, err = h.uc.ProcessingBooksGet(userID)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	fmt.Println(len(pb))

	respfmt.OK(c, pb)
}
func (h *Handler) ProcessingBooksDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, "id is not correct")
	}

	if err := h.uc.ProcessingBooksDelete(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}

func (h *Handler) BooksReviewsGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	bookID, err := strconv.Atoi(c.Request.URL.Query().Get("book_id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	reviews, err := h.uc.BooksReviewsGet(bookID, limit, offset)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, reviews)
}
func (h *Handler) BooksReviewGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	bookID, err := strconv.Atoi(c.Request.URL.Query().Get("book_id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	reviews, err := h.uc.BooksReviewGet(bookID)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, reviews)
}
func (h *Handler) BooksReviewsAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var review models.Review

	if err := c.BindJSON(&review); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	id, err := h.uc.StoreBookReview(review)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) BooksReviewsUpdate(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var review models.Review

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := c.BindJSON(&review); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	reviewID, err := h.uc.UpdateBookReview(id, review)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, map[string]interface{}{
		"id": reviewID,
	})
}
func (h *Handler) BooksReviewsDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.DeleteBookReview(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}

func (h *Handler) BooksReviewsGradesGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	grades, err := h.uc.BooksReviewsGradesGet(userID, bookID)
	if err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, grades)
}
func (h *Handler) BooksReviewsGradesAdd(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	var grades models.ReviewGrades
	if err := c.BindJSON(&grades); err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksReviewsGradesAdd(grades); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
func (h *Handler) BooksReviewsGradesDelete(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respfmt.BadRequest(c, err.Error())
		return
	}

	if err := h.uc.BooksReviewsGradesDelete(id); err != nil {
		respfmt.InternalServer(c, err.Error())
		return
	}

	respfmt.OK(c, "ok")
}
