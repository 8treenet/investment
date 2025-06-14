// Package repository generated by 'freedom new-project investment'
package repository

import (
	"investment/domain/po/custom"
	"testing"
	"time"
)

func TestCLSRepository_SaveArticle(t *testing.T) {
	unitTest := getSatUnitTest()
	unitTest.Run()

	var repo *CLSRepository
	unitTest.FetchRepository(&repo)
	err := repo.SaveArticle(&custom.ClsDepthArticle{
		ArticleID: 1,
		Ctime:     time.Now(),
		Created:   time.Now(),
		SortScore: 1,
		Title:     "ddd",
		Brief:     "aa",
		Content:   "11",
	})
	t.Log(err)
}

func TestCLSRepository_GetArticle(t *testing.T) {
	unitTest := getSatUnitTest()
	unitTest.Run()

	var repo *CLSRepository
	unitTest.FetchRepository(&repo)

	data, err := repo.GetArticle(1)
	if err != nil {
		panic(err)
	}
	jsonLog(data)
}
