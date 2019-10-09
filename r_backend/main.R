# > version
# _                           
# platform       x86_64-w64-mingw32          
# arch           x86_64                      
# os             mingw32                     
# system         x86_64, mingw32             
# status                                     
# major          3                           
# minor          5.2                         
# year           2018                        
# month          12                          
# day            20                          
# svn rev        75870                       
# language       R                           
# version.string R version 3.5.2 (2018-12-20)
# nickname       Eggshell Igloo     

if(!"quantmod" %in% installed.packages()[, 1]) {
  install.packages("quantmod")
}

library("quantmod")

# So far quantmod seems like best choice 
# return object: 
#   ticker ??
#   time 
#   forex data (which specific kind is it?)
#   NA's???

# forex data --------------------------------------------------------------

forexGetter <- function(TICKER = "USD/JPY", TO = "2014-06-02", 
                        FROM = Sys.Date()) {
  # TO, FROM: CCYY - MM - DD 
  # TICKER: CURRENCY/CROSS  
  
  # Construct call dates for date range
  time_diff <- as.integer(invisible(difftime(FROM, TO, units = "days")))
  n_months <- time_diff %/% 179
  n_rest <- time_diff %% 179
  
  call_dates <- sapply(seq_len(n_months), function(i) as.Date(FROM) - i * 179, 
                       simplify = FALSE)
  
  n <- length(call_dates)
  call_dates[n + 1] <- call_dates[n][[1]] - n_rest
  
  out <- vector("list", n + 1)
  for(i in seq(from = 2, to = n + 1)) {
    tmp_from <- call_dates[[i + 1]]
    tmp_to <- call_dates[[i]]
    
    tmp_data <- new.env() 
    
    
    if(i < 2) {
      out[[1]]
    }
  }
  
  data_out <- new.env()
#  quantmod::getFX(TICKER, FROM, TO, env = data_out)
  return(call_dates)
}

library(xts)
tick1 <- "USD/JPY"
out <- forexGetter(tick1)
Sys.Date() - 179 * (1:3)
a <- invisible(difftime(Sys.Date(), "2005-06-02", units = "days"))

sapply(seq_len(5), function(i) Sys.Date() - 179 * i, simplify = F)



# gdp data ----------------------------------------------------------------







