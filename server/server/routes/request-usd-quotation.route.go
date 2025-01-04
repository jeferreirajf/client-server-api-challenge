package route

import (
	"encoding/json"
	"net/http"

	requestUsecase "github.com/jeferreirajf/client-server-api-challenge/server/usecase/request-quotation"
)

type RequestUsdQuotationRoute struct {
	requestQuotation requestUsecase.RequestQuotationUsecase
}

func NewRequestUsdQuotationRoute(requestQuotation requestUsecase.RequestQuotationUsecase) RequestUsdQuotationRoute {
	return RequestUsdQuotationRoute{
		requestQuotation: requestQuotation,
	}
}

func (h *RequestUsdQuotationRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	inputDto := requestUsecase.RequestQuotationInputDto{}

	outputDto, err := h.requestQuotation.Execute(&inputDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(outputDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
