func restoreIpAddresses(s string) []string {
    answer := []string{}
    dfs(0,s,&answer, [][]string{}, []string{}, 0)
    return answer
}

func dfs(idx int, s string, answer *[]string,ip [][]string,block []string, num int){
    if num > 255{
        return
    }
    if len(block) == 2 && block[0] == "0"{
        return
    }
    if len(ip) > 1 && len(ip[len(ip)-2]) == 0{
        return
    }
    if len(ip) == 4{
        return
    }
    if idx == len(s){
        if len(ip) == 3{
            ip = append(ip,block)
            curr_ip := []string{}
            for _,block := range(ip){
                curr_ip = append(curr_ip,strings.Join(block, ""))
            }
            *answer = append(*answer,strings.Join(curr_ip, "."))
            fmt.Println(ip)
            ip = ip[:len(ip)-1]
        }
        return
    }
    
    item := string(s[idx])
    digit, _ := strconv.Atoi(item)
    num *=10
    num +=int(digit)
    block = append(block,item)
    dfs(idx+1,s,answer,ip,block,num)
    block = block[:len(block)-1]
    num -= int(digit)
    num /= 10

    ip = append(ip,block)
    dfs(idx+1,s,answer,ip,[]string{item},digit)
    ip = ip[:len(ip)-1]
}