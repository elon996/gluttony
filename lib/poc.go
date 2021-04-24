package lib

type Poc struct {
    Cve Cve
    Requests []Request
    Single bool
    // SendType string
}

type Cve struct {
    Id string
    Product string
}
































