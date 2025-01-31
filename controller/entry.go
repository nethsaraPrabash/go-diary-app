package controller

import (
    "go-diary-app/model"
    "go-diary-app/helper"
    "github.com/gin-gonic/gin"
    "net/http"
)

func AddEntry(context *gin.Context) {
    var input model.Entry
    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    input.UserId = user.ID

    savedEntry, err := input.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
    user, err := helper.CurrentUser(context)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}