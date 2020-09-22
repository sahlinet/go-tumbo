package handler

/*
func GetAllWorkers(db *gorm.DB, c *gin.Context) func(*gin.Context){
	employees := []model.Worker{}
	db.Find(&employees)
	//respondJSON(w, http.StatusOK, employees)
	message := "hello"
	return func(c *gin.Context) {
		c.String(http.StatusOK, message)
	}
}

func CreateWorker(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	worker := model.Worker{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&worker); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&worker).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, worker)
}

*/
