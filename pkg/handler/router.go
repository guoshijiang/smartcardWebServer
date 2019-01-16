package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.POST("/api/notice/queryNotice", httprouter.Handle(QueryNotice))
	r.POST("/api/notice/addNotic", httprouter.Handle(AddNotice))
	r.POST("/api/notice/delNotic", httprouter.Handle(DelNotice))
	r.POST("/api/notice/updateNotic", httprouter.Handle(UpdateNotice))

	r.POST("/api/news/QueryNews", httprouter.Handle(QueryNews))
	r.POST("/api/news/AddNews", httprouter.Handle(AddNews))
	r.POST("/api/news/DelNews", httprouter.Handle(DelNews))
	r.POST("/api/news/UpdateNews", httprouter.Handle(UpdateNews))


	r.POST("/api/appDownload/QueryAppDownload", httprouter.Handle(QueryAppDownload))
	r.POST("/api/appDownload/AddAppDownload", httprouter.Handle(AddAppDownload))
	r.POST("/api/appDownload/DelAppDownload", httprouter.Handle(DelAppDownload))
	r.POST("/api/appDownload/UpdateAppDownload", httprouter.Handle(UpdateAppDownload))


	r.POST("/api/partner/QueryPartner", httprouter.Handle(QueryPartner))
	r.POST("/api/partner/AddPartner", httprouter.Handle(AddPartner))
	r.POST("/api/partner/DelPartner", httprouter.Handle(DelPartner))
	r.POST("/api/partner/UpdatePartnerpartner", httprouter.Handle(UpdatePartner))


	r.POST("/api/teamIntroduce/QueryTeamIntroduce", httprouter.Handle(QueryTeamIntroduce))
	r.POST("/api/teamIntroduce/AddTeamIntroduce", httprouter.Handle(AddTeamIntroduce))
	r.POST("/api/teamIntroduce/DelTeamIntroduce", httprouter.Handle(DelTeamIntroduce))
	r.POST("/api/teamIntroduce/UpdateTeamIntroduce", httprouter.Handle(UpdateTeamIntroduce))

	return r
}
