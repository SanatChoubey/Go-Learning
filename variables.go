package main

import "fmt"

func main() {
	students := [...]string{"bhavesh", "sanat", "sambhav", "abc"}
	fmt.Printf("%v \n", students)
}
import React from 'react';

const FilterComponent = () => {
	const [list, setList] = React.useState(['apple', 'orange', 'peach']);
	const [enteredText, setEnteredText] = React.useState('');

	return (
		<div>
		{
			list.map((item, i) => {
				return <div key={`${item} ${i}`}>{item}</div>
			})
		}
		<input type="text" placeholder="enter text" onChange={(e) => {setEnteredText(e.target.value)}} />
		<button onClick={() => {
			const memoArray = [...list];
			const indexofWord = list.indeOf(enteredText)
			if(indexofWord > -1) {
				memoArray.splice(indexofWord, 1)
				setList(memoArray)
			}
		}} >filter out</button>
		</div>
		

	)
}