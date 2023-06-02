package main

type Transaction struct{
	IDs []int `json:"ids" binding:"required"`
	Total int `json:"total" binding:"required"`
	PaidBy int `json:"paidby" binding:"required"`
	Shares []int `json:"shares"`
}

type Person struct{
	Name string `json:"name" binding:"required"`
	ID int `json:"id"`
}