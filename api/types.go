/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package api

const (
	LeetcodeCn = "https://leetcode.cn"
	Leetcode   = "https://leetcode.com"
)

// QuestionQuery graphql query
const QuestionQuery = `
	query questionOfToday {
    todayRecord  {
        date
        question {
            acRate
            difficulty
            title
            titleSlug
            topicTags {
                name
                id
            }
        }
    }
}
`

// QuestionTodayResp 注意与官网直接 restful api 请求返回的少一个 data 字段嵌套
type QuestionTodayResp struct {
	TodayRecord []struct {
		Date     string `json:"date"`
		Question struct {
			Difficulty string  `json:"difficulty"`
			Title      string  `json:"title"`
			TitleSlug  string  `json:"titleSlug"`
			AcRate     float64 `json:"acRate"`
			TopicTags  []struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"todayRecord"`
}
