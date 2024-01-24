import React, { useState, useEffect } from 'react'
import {ethers} from 'ethers'
import {abi} from '../../../artifacts/contracts/Vote.sol/Vote.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"
import { setuid } from 'process'


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;


export default function CommitVote(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [uid, setUid] = useState(0)
  var [prompt, setPrompt] = useState<string>("")
  var [type, setType] = useState(0)

  async function commitVote(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.BrowserProvider(window.ethereum)
    const signer = await provider.getSigner()
    const Vote:Contract = new ethers.Contract(addressContract, abi, signer)

    if (type == 0) {

    Vote.CommitChoiceFreePromt(uid, prompt)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("vote receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     } else if (type == 1) {
      Vote.CommitChoiceENSValid(uid, prompt)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("vote receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
    } else if (type == 2) {
      Vote.CommitChoice_ENS_and_T3P(uid, prompt)
      .then((tr: TransactionResponse) => {
         console.log(`TransactionResponse TX hash: ${tr.hash}`)
         tr.wait().then((receipt:TransactionReceipt) => {console.log("vote receipt", receipt)})
         })
          .catch((e:Error) => console.log(e))
    }
  }
  
  return (
    <form onSubmit={commitVote}>
      <span>UID</span>
      <input id="UID" type="number" required  onChange={(e) => setUid(parseInt(e.target.value))} value={uid} />
      <span>prompt</span>
      <input id="operator" type="text" required  onChange={(e) => setPrompt(e.target.value)} value={prompt} />
      <span>votetype</span>
      <input id="votetype" type="number" required  onChange={(e) => setType(parseInt(e.target.value))} value={type} />
      <button type="submit" >Submit Vote</button>
   
    </form>
  )
}