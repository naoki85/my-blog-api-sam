package repository

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.ResistMock("^SELECT (.+) FROM recommended_books .*", []string{"id", "link", "image_url", "button_url"})
	repo := RecommendedBookRepository{
		SqlHandler: mockSqlHandler,
	}
	recommendedBooks, err := repo.All(4)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if len(recommendedBooks) != 4 {
		t.Fatalf("Fail expected: 4, got: %v", len(recommendedBooks))
	}
	if recommendedBooks[0].Link != "http://naoki85.test" {
		t.Fatalf("Fail expected: http://naoki85.test, got: %v", recommendedBooks[0].Link)
	}
}

func TestShouldCreateRecommendedBook(t *testing.T) {
	mockSqlHandler, _ := NewMockSqlHandler()
	mockSqlHandler.Mock.ExpectExec("INSERT INTO recommended_books").
		WillReturnResult(sqlmock.NewResult(1, 1))
	repo := RecommendedBookRepository{
		SqlHandler: mockSqlHandler,
	}
	params := RecommendedBookCreateParams{
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}
	err := repo.Create(params)
	if err != nil {
		t.Fatalf("Cannot create recommended_book: %s", err)
	}
	if err := mockSqlHandler.Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
