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

package vo

//ErrorVo This error object contain code, message and its explanation.
type ErrorVo struct {
	Code             int    `json:"code,omitempty"`
	Message          string `json:"message,omitempty"`
	ErrorExplanation string `json:"errorExplanation,omitempty"`
}
