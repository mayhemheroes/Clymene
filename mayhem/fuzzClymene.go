package fuzzClymene

import "strconv"
import "github.com/Clymene-project/Clymene/cmd/promtail/app/logentry/stages"
import "github.com/Clymene-project/Clymene/pkg/lokiutil"
import "github.com/Clymene-project/Clymene/model/labels"
import "github.com/Clymene-project/Clymene/cmd/promtail/app/logentry/logql"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            var test stages.Packed
            test.UnmarshalJSON(bytes)
            return 0

        case 1:
            var strArr = make([]string, len(bytes))
            for i, byte := range bytes {

                strArr[i] = string(byte)
            }

            util.MergeStringLists(strArr)
            return 0

        case 2:
            content := string(bytes)
            util.SnakeCase(content)
            return 0

        case 3:
            content := string(bytes)
            list := []string{"mayhem", "fuzz"}
            util.StringsContain(list, content)
            return 0

        case 4:
            content := string(bytes)
            labels.NewFastRegexMatcher(content)
            return 0

        case 5:
            content := string(bytes)
            logql.ParseExpr(content)
            return 0

        default:
            content := string(bytes)
            logql.ParseMatchers(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}