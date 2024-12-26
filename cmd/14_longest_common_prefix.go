package cmd

func longestCommonPrefix(strs []string) string {
    var longCommon string

    shortestWord := strs[0]
    for _, v := range(strs) {
        if len(v) < len(shortestWord) {
            shortestWord = v
        }
    }
    var running = ""
    for _, v := range(shortestWord) {
        running += string(v)
        if (checkAllWordsConainStr(running, strs)) {
            if (len(running) > len(longCommon)) {
                longCommon = running
            }
        }
    }

    return longCommon   
}

func checkAllWordsConainStr ( s string, words []string) bool {
    for _, w := range(words) {
        if (w[:len(s)] != s) {
            return false
        }
    }
    return true
}