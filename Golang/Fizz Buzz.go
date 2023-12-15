import "strconv"
func fizzBuzz(n int) []string {
    ans := make([]string, 0)
    for i:=1; i <=n ;i++ {
        if (i % 3 == 0 && i % 5 == 0){
            ans = append(ans, "FizzBuzz")
        } else if (i % 5 == 0) {
            ans = append(ans, "Buzz")
        } else if (i % 3 == 0) {
            ans = append(ans, "Fizz")
        } else {
            ans = append(ans, strconv.Itoa(i))
        }
    }

    return ans
}