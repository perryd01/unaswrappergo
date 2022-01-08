# Unas API wrapper lib for Golang
[![Coverage Status](https://coveralls.io/repos/github/perryd01/unaswrappergo/badge.svg?branch=main)](https://coveralls.io/github/perryd01/unaswrappergo?branch=main) [![Go Reference](https://pkg.go.dev/badge/github.com/perryd01/unaswrappergo.svg)](https://pkg.go.dev/github.com/perryd01/unaswrappergo)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/perrydlol)

- [Unas API wrapper lib for Golang](#unas-api-wrapper-lib-for-golang)
  - [Contributions](#contributions)
  - [Maintenance](#maintenance)
  - [Code documentation](#code-documentation)
  - [Usable API requests](#usable-api-requests)


This is an unofficial [API](https://unas.hu/tudastar/api) wrapper lib for [Unas webshop](https://unas.hu/) written for Golang. The main goal is to make programming stuff using Unas API more convinient.
## Contributions
Big thanks to [@ondrejoa](https://github.com/ondrejoa) for helping me in this.

## Maintenance
The largest part of the codebase was written in 2021 summer. We cannot guarantee the correct operation in case of an API update. **Feel free to contribute and keep the code up to date.**

## Code documentation
Unas API has an often poor, so we can't figure out some minor stuff either :confused:

## Usable API requests

| Endpoint            | Support              |
|---------------------|----------------------|
| AuthwithAPIKey      | :heavy_check_mark:   |
| AuthwithPass        | not gonna support it |
| GetNewsletter       | :question:           |
| SetNewsletter       | :question:           |
| GetProduct          | :x:                  |
| SetProduct          | :x:                  |
| GetProductDB        | :x:                  |
| SetProductDB        | :x:                  |
| GetOrder            | :x:                  |
| SetOrder            | :x:                  |
| GetStock            | :x:                  |
| SetStock            | :x:                  |
| GetCategory         | :heavy_check_mark:   |
| SetCategory         | :heavy_check_mark:   |
| GetCustomer         | :x:                  |
| SetCustomer         | :x:                  |
| CheckCustomer       | :question:           |
| GetScriptTag        | :x:                  |
| SetScriptTag        | :x:                  |
| GetPage             | :x:                  |
| SetPage             | :x:                  |
| GetPageContent      | :x:                  |
| SetPageContent      | :x:                  |
| GetStorage          | :heavy_check_mark:   |
| SetStorage          | :heavy_check_mark:   |
| GetProductParameter | :x:                  |
| SetProductParameter | :x:                  |
| GetAutomatism       | :x:                  |
| SetAutomatism       | :x:                  |