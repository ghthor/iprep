[source](https://www.interviewcake.com/question/javascript/stock-price)

## My Solution

Starting at the end of list of prices, we walk backwards storing the highest
value we find. This should handle the cases where the prices stayed the same
all day(possible profit == 0) and where the prices dropped all day(possible
profit == negitive).
