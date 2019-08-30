// package apihelper
// Copyright (c) 2019-2099, PIXEOM, INC. (http://www.pixeom.com)
// All rights reserved.  PIXEOM HIGHLY CONFIDENTIAL
// THIS IS PROPRIETARY SOFTWARE OWNED BY PIXEOM, INC.
// THE SOURCE CODE IS PROVIDED ONLY UNDER NDA AND SHOULD BE HELD
// IN STRICTEST CONFIDENCE.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL PIXEOM, INC. OR ITS CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package apihelper

import (
	"encoding/json"
	"license/api/vo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateResponseWithErrExplanation This method create response for ERROR with explanationMessage.
//@param message : error message
//@param explanationMessage : error message explanation
//@param code : error code
func CreateResponseWithErrExplanation(w *gin.Context, message, explanationMessage string, code int) {

	errorObject := vo.ErrorVo{}
	errorObject.Code = code
	errorObject.Message = message
	errorObject.ErrorExplanation = explanationMessage
	CreateResponseWithError(w, errorObject)

}

func CreateResponseWithError(w *gin.Context, errorObject vo.ErrorVo) {
	errStr, err := json.Marshal(errorObject)
	if err != nil {
		log.Println(err)
		return
	}
	http.Error(w.Writer, string(errStr), errorObject.Code)
}

//CreateResponse This method create response for ERROR.
func CreateResponse(w *gin.Context, payload interface{}) {
	w.JSON(200, payload)
}
