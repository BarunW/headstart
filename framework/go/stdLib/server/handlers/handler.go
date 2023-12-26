package handlers

/*
**
======================================================
Replace the storage type with actual db driver type
=====================================================
**
*/
type Handler struct {
	db any
}

func NewHandler(storage any) *Handler {
	return &Handler{
		db: storage,
	}
}

type AnnoyingMsg struct {
	msg    string
	status int
}

func (h *Handler) jsonResponseWriter(w http.ResponseWriter, status int, value any) error {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return encoder.Encode(AnnoyingMsg{msg: "hello world", status: 100})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.jsonResponseWriter(w, http.StatusOK, "happu")
	if err != nil {
		fmt.Println("error")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
