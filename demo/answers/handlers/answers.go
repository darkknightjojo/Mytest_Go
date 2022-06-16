package handlers

import (
	"answers/pkg/database"
	"answers/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func SaveAnswer(c *gin.Context) {
	ip := c.Request.Header.Get("Referer")
	session := sessions.Default(c)
	u := c.PostForm("user")
	i := c.PostForm("id")
	body := c.PostForm("body")
	answerUserId, _ := strconv.Atoi(u)
	questionUserId, _ := strconv.Atoi(i)
	answer := models.Answer{
		UserID:     answerUserId,
		QuestionID: questionUserId,
		Body:       body,
	}

	// 保存答案
	database.DB.Save(&answer)
	// 更新问题中答案数
	database.DB.Exec("UPDATE questions SET answer_count = answer_count + 1 WHERE questions.id = ?", questionUserId)

	_ = session.Save()
	c.Redirect(http.StatusFound, ip)
}

func AcceptAnswer(c *gin.Context) {
	var answer []models.Answer
	var question []models.Question
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)

	// 将答案设置为被接收状态
	database.DB.Model(&answer).
		Where("answer.id = ?", answerId).
		UpdateColumn("is_accept_answer", gorm.Expr("is_accepted_answer + ?", 1))

	// 设置问题状态为已接收答案
	database.DB.Model(&question).
		Where("question.id = ?", questionId).
		UpdateColumn("accepted_answer", gorm.Expr("accepted_answer + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func EditAnswer(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	user := session.Get("user")
	answer := models.Answer{}
	var users []models.User

	var answerUserID int
	database.DB.Find(&users)
	for _, v := range users {
		if v.Username == user {
			answerUserID = v.Id
		}
	}

	database.DB.Find(&answer, id)

	c.HTML(http.StatusOK, "answeredit.tmpl.html",
		gin.H{"answer": answer,
			"user":         user,
			"answerUserID": answerUserID,
		})
}

func AnswerLikes(c *gin.Context) {
	var answer []models.Answer
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&answer).
		Where("answers.id = ?", answerId).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func AnswerDislikes(c *gin.Context) {
	var answer []models.Answer
	id := c.PostForm("qid")
	ans := c.PostForm("aid")
	answerId, _ := strconv.Atoi(ans)
	questionId, _ := strconv.Atoi(id)
	database.DB.Model(&answer).
		Where("answers.id = ?", answerId).
		UpdateColumn("dis_likes", gorm.Expr("dis_likes + ?", 1))

	t := strconv.Itoa(questionId)
	c.Redirect(http.StatusFound, "/show/"+t)
}

func UpdateAnswer(c *gin.Context) {
	id := c.Param("id")
	session := sessions.Default(c)
	body := c.PostForm("body")
	answer := models.Answer{}

	database.DB.Model(&answer).
		Where("id = ?", id).
		Update("body", body)

	// TODO session.Save()是做什么的？
	_ = session.Save()
	c.Redirect(http.StatusFound, "/")
}

func AnswerDelete(c *gin.Context) {
	id := c.Param("id")
	ip := c.Request.Header.Get("Referer")
	answerId, _ := strconv.Atoi(id)
	var answers []models.Answer
	database.DB.Delete(&answers, answerId)
	c.Redirect(http.StatusFound, ip)
}
