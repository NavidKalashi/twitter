import "github.com/gin-gonic/gin"


type Server struct {
	tweetController *TweetController
}

func NewServer(tweetController *TweetController) *Server {	
	return &Server{tweetController: tweetController}
}


func (s *Server) AddRoutes() {
}


func (s *Server) Start() {
}
