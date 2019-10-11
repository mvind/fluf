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

library("quantmod")""

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
  
  env_name <- stringr::str_replace_all(TICKER, "/", "")
  cat("envname: ", env_name, "\n")
  for(i in seq(from = 2, to = n + 1)) {
    tmp_from <- call_dates[[i - 1]]
    tmp_to <- call_dates[[i]]
    
    cat(tmp_from, tmp_to)
    stop()
    tmp_data <- new.env() 
    
    quantmod::getFX(TICKER, from = tmp_from, to = tmp_to, env = tmp_data)
    
    #get(env_name, envir = tmp_data)
    out[[i - 1]] <- get(env_name, envir = tmp_data)
    cat(i - 1, " out of ", n, " \n")
    rm(tmp_data)
  
  }
  
  return(out)
}

res <- forexGetter(TO = "2018/06/02")

library(xts)
tick1 <- "USD/JPY"
test_env <- new.env()
getFX(tick1, env = test_env)


. <-get("USDJPY", test_env)





# gdp data ----------------------------------------------------------------







