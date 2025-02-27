package daily

// TextEditor 2296. 设计一个文本编辑器
type TextEditor struct {
	left, right []byte
}

//func Constructor() TextEditor {
//	return TextEditor{}
//}

func (t *TextEditor) AddText(text string) {
	t.left = append(t.left, text...)
}

func (t *TextEditor) DeleteText(k int) int {
	k = min(k, len(t.left))
	t.left = t.left[:len(t.left)-k]
	return k
}

func (t *TextEditor) text() string {
	return string(t.left[max(0, len(t.left)-10):])
}

func (t *TextEditor) CursorLeft(k int) string {
	for k > 0 && len(t.left) > 0 {
		t.right = append(t.right, t.left[len(t.left)-1])
		t.left = t.left[:len(t.left)-1]
		k--
	}
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	for k > 0 && len(t.right) > 0 {
		t.left = append(t.left, t.right[len(t.right)-1])
		t.right = t.right[:len(t.right)-1]
		k--
	}
	return t.text()
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
