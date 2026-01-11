# Guide to Querying Wolfram Alpha with Natural Language

Wolfram Alpha is a computational knowledge engine that allows users to input queries in natural language (plain English or similar phrasing) to compute answers across a wide range of topics, from mathematics and science to history, finance, and everyday facts. It interprets your words to perform calculations, generate visualizations, and provide data-driven insights. No programming knowledge is required—just type as you would describe the problem conversationally.

## General Tips for Natural Language Queries
- **Be descriptive but concise**: Use everyday language to describe what you want. For example, instead of complex formulas, say "integrate x^2 from 0 to 1" or "what is the distance from Earth to Mars?"
- **Specify details**: Include units, parameters, or contexts to refine results (e.g., "convert 5 feet to meters" or "population of Tokyo in 2020").
- **Handle ambiguity**: If results aren't what you expect, rephrase for clarity (e.g., add "exactly" for precise counts or "probability of" for chance-based queries).
- **Use math notation optionally**: Natural language works best, but you can mix in symbols like +, -, *, /, ^ for exponents, or integrals.
- **Explore step-by-step**: Wolfram Alpha often provides step-by-step solutions; append "step by step" to your query if needed.
- **Access it online**: Go to [wolframalpha.com](https://www.wolframalpha.com) and type your query in the search bar.
- **Limitations**: It excels at factual, computational queries but may not handle highly subjective or real-time breaking news well.

## Broad Examples Across Categories
Wolfram Alpha covers diverse fields. Here are examples from various categories to illustrate its versatility. Each includes a sample natural language query and a brief description of what it computes.

### Mathematics (General)
- **Query**: "solve x^2 + 3x - 4 = 0"  
  Solves the quadratic equation and provides roots, graph, and alternate forms.
- **Query**: "plot sin(x) from -pi to pi"  
  Generates a plot of the sine function over the specified interval.
- **Query**: "factor 12x^2 - 18x + 6"  
  Factors the polynomial expression.
- **Query**: "derivative of e^x * sin(x)"  
  Computes the derivative and shows steps.

### Science and Physics
- **Query**: "force of gravity on 5kg mass on Earth"  
  Calculates gravitational force using standard Earth values.
- **Query**: "chemical formula for water"  
  Provides molecular structure, properties, and related data.
- **Query**: "speed of light in vacuum"  
  Returns the value with units and conversions.
- **Query**: "orbit of Mars"  
  Shows orbital parameters, diagrams, and current position.

### Chemistry and Biology
- **Query**: "molecular weight of caffeine"  
  Computes the molar mass and provides structural formula.
- **Query**: "DNA sequence for human insulin gene"  
  Retrieves genetic sequence data and related biology info.
- **Query**: "pH of 0.1M HCl"  
  Calculates the pH value for the acid solution.

### History and Society
- **Query**: "when was the moon landing"  
  Provides date, details, and historical context for Apollo 11.
- **Query**: "population of France in 1800"  
  Gives historical population data with trends.
- **Query**: "GDP of USA vs China"  
  Compares gross domestic product with charts and projections.

### Everyday Life and Finance
- **Query**: "weather in Paris tomorrow"  
  Forecasts temperature, precipitation, and conditions (note: uses current data).
- **Query**: "mortgage payment for $300000 at 4% over 30 years"  
  Calculates monthly payments, total interest, and amortization.
- **Query**: "calories in a banana"  
  Provides nutritional breakdown.
- **Query**: "stock price of AAPL last year"  
  Shows historical stock data with charts.

### Geography and Units
- **Query**: "distance from New York to London"  
  Computes straight-line and travel distances with maps.
- **Query**: "convert 100 km/h to mph"  
  Performs unit conversion with equivalents.

## Special Focus: Probabilities and Statistics
Wolfram Alpha is particularly powerful for probability and statistics queries, handling distributions, events, and data analysis in natural language. It can compute exact probabilities, expected values, confidence intervals, and more. Below are expanded examples, including combinatorial problems, distributions, hypothesis testing, and data summaries. I've included the user's suggested dice example and similar ones for emphasis.

### Basic Probability Events
- **Query**: "throw three die get two sixes" (or more precisely: "probability of exactly two sixes in three dice rolls")  
  Computes the binomial probability of getting exactly two sixes when rolling three fair six-sided dice (result: about 0.0694 or 3/43).
- **Query**: "probability of heads on a coin flip"  
  Returns 1/2, with explanations of fair coin assumptions.
- **Query**: "probability of rolling a 7 with two dice"  
  Lists all outcomes and computes 1/6.
- **Query**: "chance of drawing an ace from a deck"  
  Calculates 4/52 or about 0.0769 for a standard deck.

### Combinatorial and Game Probabilities
- **Query**: "poker full house"  
  Computes the probability of being dealt a full house in five-card poker (about 0.00144).
- **Query**: "lottery odds for 6/49"  
  Calculates the probability of winning a standard 6/49 lottery (1 in 13,983,816).
- **Query**: "birthday paradox 50 people"  
  Computes the probability that at least two share a birthday in a group of 50 (about 0.97).
- **Query**: "probability 5 people born on same day of week"  
  Calculates the chance all five share the same weekday (about 0.000045).
- **Query**: "bet the corner at roulette"  
  Analyzes win probability (about 0.108) and expected value for betting on four numbers.

### Bernoulli Trials and Streaks
- **Query**: "probability of 8 successes in 14 trials with p=0.6"  
  Uses binomial distribution to compute the exact probability (about 0.157).
- **Query**: "streak of 12 successes in 40 trials"  
  Estimates the probability of at least one streak of 12 heads in 40 coin flips.
- **Query**: "number of trials until 15th success"  
  Computes expected value for geometric distribution (e.g., with p=0.5, it's 30 trials).
- **Query**: "32 coin tosses"  
  Provides probabilities for various outcomes like all heads or balanced results.

### Probability Formulas and Conditional Events
- **Query**: "conditional probability formula"  
  Explains P(A|B) = P(A and B)/P(B) with examples.
- **Query**: "probability of union of three events"  
  Computes P(A or B or C) using inclusion-exclusion.
- **Query**: "probability of a complement"  
  Shows 1 - P(event) with illustrative cases.

### Probability Distributions
- **Query**: "Poisson distribution lambda=5"  
  Plots the distribution and computes probabilities like P(X=3).
- **Query**: "normal distribution mean=0 sd=2"  
  Provides PDF, CDF, and quantiles for the specified normal distribution.
- **Query**: "standard deviation of Student t 17 degrees of freedom"  
  Computes variance and SD (about 1.069 for SD).

### Descriptive Statistics
- **Query**: "mean of {21.3, 38.4, 12.7, 41.6}"  
  Calculates the arithmetic mean (about 28.5).
- **Query**: "kurtosis of {21.3, 38.4, 12.7, 41.6}"  
  Computes kurtosis to measure tailedness (about -1.36).
- **Query**: "descriptive statistics for {25, 35, 10, 17, 29, 14, 21, 31}"  
  Returns mean, median, mode, range, variance, and more.

### Inferential Statistics and Hypothesis Testing
- **Query**: "t-interval xbar=4.15 s=0.32 n=100"  
  Computes a 95% confidence interval for the population mean.
- **Query**: "z-test for mean mu=50 xbar=52 s=5 n=30"  
  Performs a hypothesis test and gives p-value.
- **Query**: "sample size for binomial p=0.5 margin=0.03"  
  Estimates required n for a confidence interval.

### Random Variables and Expectations
- **Query**: "EV[3X^4 - 7] for X~Poisson(7.3)"  
  Computes the expected value using the distribution.
- **Query**: "P[-1.2 < X < 2.3] for X~student t 12 dof"  
  Calculates the cumulative probability (about 0.877).

### Data Fitting and Modeling
- **Query**: "linear fit {1.3,2.2}, {2.1,5.8}, {3.7,10.2}, {4.2,11.8}"  
  Fits a line and provides equation, R-squared, and plot.
- **Query**: "exponential fit 0.783,0.552,0.383,0.245,0.165,0.097"  
  Models exponential decay with parameters.

## Special Focus: Finance, Investing, and Trading
Wolfram Alpha excels at financial computations, pulling real-time and historical data on stocks, currencies, bonds, economic indicators, and more. It can calculate ratios, model investments, forecast trends, and handle complex trading scenarios using natural language. Queries often start with terms like "stock price of," "financial ratios for," or "return on investment for." Below are examples, emphasizing practical applications for investors and traders.

### Stock Market and Securities
- **Query**: "AAPL stock price"  
  Retrieves the current price of Apple stock, including intraday chart, volume, and key stats like open/high/low/close.
- **Query**: "historical stock price of TSLA from 2020 to 2025"  
  Provides a historical price chart, returns, volatility, and downloadable data for Tesla over the specified period.
- **Query**: "compare stock performance of AMZN vs MSFT last year"  
  Compares total returns, price charts, and metrics like beta for Amazon and Microsoft.
- **Query**: "dividend yield of KO"  
  Calculates the current dividend yield for Coca-Cola, with payout history and ex-dividend dates.
- **Query**: "market capitalization of Google"  
  Computes Alphabet's (GOOGL) market cap, with breakdowns by shares outstanding and comparisons to peers.

### Financial Ratios and Analysis
- **Query**: "P/E ratio of NFLX"  
  Returns Netflix's price-to-earnings ratio, forward P/E, and industry comparisons.
- **Query**: "ROE for JPMorgan Chase"  
  Calculates return on equity, with formulas, historical trends, and DuPont analysis.
- **Query**: "debt to equity ratio of Ford Motor Company"  
  Provides the D/E ratio, balance sheet excerpts, and solvency insights.
- **Query**: "beta of Bitcoin vs S&P 500"  
  Computes the beta coefficient measuring Bitcoin's volatility relative to the S&P 500 index.
- **Query**: "Sharpe ratio for a portfolio of 50% AAPL 50% bonds over 5 years"  
  Estimates the risk-adjusted return using historical data, assuming a risk-free rate.

### Investing and Portfolio Management
- **Query**: "compound interest on $10000 at 5% for 10 years"  
  Calculates future value with compounding, including annual breakdowns and graphs.
- **Query**: "return on investment for $5000 in gold from 2010 to now"  
  Computes ROI, annualized return, and inflation-adjusted performance for gold.
- **Query**: "efficient frontier for stocks AAPL MSFT GOOG"  
  Generates a Markowitz efficient frontier plot for a portfolio of these stocks based on historical returns.
- **Query**: "Monte Carlo simulation for retirement savings $200k initial 7% return 30 years"  
  Runs simulations to estimate portfolio growth, with probability distributions for outcomes.
- **Query**: "diversification benefits of adding Bitcoin to S&P 500 portfolio"  
  Analyzes correlation, variance reduction, and optimized allocations.

### Trading and Options
- **Query**: "Black-Scholes price for AAPL call option strike 150 expiration in 3 months"  
  Computes the theoretical option price using implied volatility and other Greeks (delta, gamma, etc.).
- **Query**: "implied volatility of SPY options"  
  Returns current IV for S&P 500 ETF options, with surfaces and historical charts.
- **Query**: "technical analysis of EUR/USD forex pair"  
  Provides charts with moving averages, RSI, MACD, and other indicators for the euro-dollar exchange rate.
- **Query**: "futures price of crude oil"  
  Shows current WTI or Brent crude futures prices, with contract details and settlement history.
- **Query**: "arbitrage opportunity between Bitcoin on Binance vs Coinbase"  
  Compares spot prices across exchanges (if data available) and highlights potential spreads.

### Currencies and Crypto
- **Query**: "convert 1000 EUR to USD"  
  Performs real-time currency conversion with exchange rates and historical trends.
- **Query**: "Bitcoin price in USD last month"  
  Charts Bitcoin's price, with volume, market cap, and halving event annotations.
- **Query**: "Ethereum vs Bitcoin correlation"  
  Computes the correlation coefficient over time, with scatter plots and rolling windows.
- **Query**: "inflation-adjusted value of $100 from 2000 to now"  
  Adjusts for US CPI inflation, showing purchasing power and rate history.

### Economic Indicators and Macroeconomics
- **Query**: "US unemployment rate"  
  Returns the current rate, with charts, forecasts, and state breakdowns.
- **Query**: "GDP growth of China vs USA last 10 years"  
  Compares annual GDP growth rates, with nominal vs real figures and projections.
- **Query**: "interest rate on 10-year Treasury bond"  
  Provides yield curve data, historical yields, and inflation expectations.
- **Query**: "consumer price index for UK"  
  Shows CPI trends, components (e.g., food, energy), and year-over-year changes.
- **Query**: "federal funds rate history"  
  Charts the Fed's target rate, with economic context like recessions.

These queries leverage Wolfram Alpha's integration with financial databases for accurate, up-to-date results. For real-time trading, append "current" or specify dates. If modeling custom scenarios, include assumptions like "assuming 2% inflation" to refine outputs. Experiment with combinations for deeper insights.
