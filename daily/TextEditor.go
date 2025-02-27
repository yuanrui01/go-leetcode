package daily

// TextEditor 2296. 设计一个文本编辑器
type TextEditor struct {
	s      []byte
	cursor int
}

//func Constructor() TextEditor {
//	return TextEditor{make([]byte, 0, 1024), 0}
//}

func (this *TextEditor) AddText(text string) {
	bytes := []byte(text)
	this.s = append(this.s[:this.cursor], append(bytes, this.s[this.cursor:]...)...)
	this.cursor += len(text)
}

func (this *TextEditor) DeleteText(k int) int {
	cnt := min(k, this.cursor)
	if cnt > 0 {
		this.s = append(this.s[:this.cursor-cnt], this.s[this.cursor:]...)
	}
	this.cursor -= cnt
	return cnt
}

func (this *TextEditor) CursorLeft(k int) string {
	this.cursor -= min(k, this.cursor)
	return string(this.s[max(0, this.cursor-10):this.cursor])
}

func (this *TextEditor) CursorRight(k int) string {
	this.cursor = min(this.cursor+k, len(this.s))
	return string(this.s[max(0, this.cursor-10):this.cursor])
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
