"use client"


import React from 'react'

const ChatHistory = ({messages}:any) => {
  const chat = messages?.map((body: string | number, index: React.Key | null | undefined)=>{
    
    return <p key={index} className=''>{body}</p>
  }
  )

console.log(messages)

  return (
    <div className='flex flex-col justify-center items-center m-10 text-lg'>
      <h2 className='border-2 border-red-600 text-4xl'>
        ChatHistory
      </h2>
        {chat}

    </div>
  )
}

export default ChatHistory