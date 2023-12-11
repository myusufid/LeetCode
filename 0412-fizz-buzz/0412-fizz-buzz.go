func fizzBuzz(n int) []string {
   var answer []string
   for i:=1;i<=n;i++{
     answer = append(answer, findFizzBuzz(i))
   }
   return answer
}

func findFizzBuzz(i int) string {
  switch {
  case i % 15 == 0:
    return "FizzBuzz"   
  case i % 5 == 0:
    return "Buzz"   
  case i % 3 == 0:
    return "Fizz"
  default:
    return strconv.Itoa(i)   
  }
}