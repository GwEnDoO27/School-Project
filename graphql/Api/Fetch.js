//Query for fetching user data login and transactions
export const query = `
{
  user {
    id
    login
    attrs
    auditRatio
    totalUp
    totalDown
    xps {
      amount
      path
    }
  }
  transaction {
    type
    amount
  }
}
`

//Query for fetching user data xps
export const query2 = `
{
  user {
    xps(order_by: { amount: desc }, limit: 3) {
      amount
      originEventId
      path
    }
  }
}
`

//Query for fetching user data only transactions xps
export const query3 = `
{
    transaction(where: {type : {_eq: "xp"}}){
      amount 
    	path 
      createdAt
    }
}
`

