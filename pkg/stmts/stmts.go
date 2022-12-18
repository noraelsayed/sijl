package stmts

const CheckUser string = `SELECT COUNT(*)
FROM SIJL.USERS
WHERE username = @username`
const UsernameConvention string = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
