import Head from "next/head";
import CreateNewVote from 'components/createNewVote'
import CommitVote from 'components/commitVote'
import { useState} from 'react'
import {ethers} from 'ethers'


declare let window: any;


export default function Home() {
  const [currentAccount, setCurrentAccount] = useState<string | undefined>()
 
  const onClickConnect = async () => {
    //client side code
    if(!window.ethereum) {
      console.log("please install MetaMask")
      return
    }
  
    const accts = await window.ethereum.request({
      method: "eth_requestAccounts"
 })
 if(accts.length>0) setCurrentAccount(accts[0])
 
}

  const onClickDisconnect = () => {
    setCurrentAccount(undefined)
  }



  return (
    <>
      <Head>
        <title>Пыпыщ</title>
      </Head>
      {currentAccount  
          ? <button type="button" onClick={onClickDisconnect}>
                Account:{currentAccount}
            </button>
          : <button type="button" onClick={onClickConnect}>
                  Connect MetaMask
              </button>
        }
     <div>Create New Vote</div>
     <CreateNewVote 
            addressContract='0xae92476c90cbA5C8c3928FFDd0CB2fF6231f4E37'
            currentAccount={currentAccount}
          />
      <div>Commit Vote</div>
      <CommitVote 
            addressContract='0xae92476c90cbA5C8c3928FFDd0CB2fF6231f4E37'
            currentAccount={currentAccount}
          />
    </>
  );
}
