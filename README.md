# Sainsbury's Go Conference Test 2017

## Subs-price Promise

Sainsbury's Chop Chop service now carries a "subs-price promise": if an item that is not part of an offer is not available and we substitute it for something more expensive, you pay the price of the original item instead.

For example:
* Alice wants to buy cookies for £1.00, but these get subsituted for cookies that cost £1.50, she will only pay £1.00.
* Bob buys cookies for £1.00 that get substituted for cookies that cost 75p, so he only pays 75p.
* Eve buys cookies _that are part of a promotion_ for £1.00, but these get substituted for cookies that cost £1.50. Eve pays £1.50. 

## The Task

Complete the GetTotal() function, which has the following signature:

```go
GetTotal(items []Item, offers []int) int
```

where offers is a list of SKUs and Item is a struct in the following format:

```go
type Item struct {
    Sku         int
    Title       string
    Price       int
    Substituted bool
    SubPrice    int
}
```

To finish the task, all unit tests must pass when `ginkgo` is run