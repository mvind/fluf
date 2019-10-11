
# Construct Test data because Oanda doesnt work right now..
library(xts)
library(lubridate)

getData <- function(){
  vals <- rnorm(100, 50, 10)
  dates <- Sys.Date() - lubridate::days(1:100)
  out <- list("USDJPY",vals, dates)
  return(out)
}

getData()