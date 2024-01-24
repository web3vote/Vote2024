import React, { useState, useEffect } from 'react'
import {ethers} from 'ethers'
import {abi} from '../../../artifacts/contracts/Vote.sol/Vote.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;


export default function CreateNewVote(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [org, setOrg] = useState<string>("")
  var [oper, setOper] = useState<string>("")
  var [time, setTime] = useState(0)
  var [hours, setHours] = useState(0)
  var [passport, setPassport] = useState(0)
  var [type, setType] = useState(0)

  async function createVote(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.BrowserProvider(window.ethereum)
    const signer = await provider.getSigner()
    const Vote:Contract = new ethers.Contract(addressContract, abi, signer)

    Vote.createNewVote(org, oper, time, hours, passport, type)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("vote receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }
  return (
    <form onSubmit={createVote}>
   
      <span>organiser</span>
      <input id="organiser" type="text" required  onChange={(e) => setOrg(e.target.value)} value={org} />
      <span>operator</span>
      <input id="operator" type="text" required  onChange={(e) => setOper(e.target.value)} value={oper} />
      <span>timestamp</span>
      <input id="timestamp" type="number" required  onChange={(e) => setTime(parseInt(e.target.value))} value={time} />
      <span>hours</span>
      <input id="hours" type="number" required onChange={(e) => setHours(parseInt(e.target.value))} value={hours}  />
      <span>passport type</span>
      <input id="passport" type="number" required onChange={(e) => setPassport(parseInt(e.target.value))} value={passport}  />
      <span>votetype</span>
      <input id="votetype" type="number" required  onChange={(e) => setType(parseInt(e.target.value))} value={type} />



      <button type="submit" >Create Vote</button>
   
    </form>
  )
}