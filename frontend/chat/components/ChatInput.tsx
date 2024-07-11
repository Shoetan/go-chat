"use client"

import { sendMsg } from '@/api/api'
import React from 'react'

const ChatInput = () => {
  const [chat, setChat] = React.useState<string>("")

  const sendChat =(e: { key: string }) =>{
    if(e.key === 'Enter'){
      sendMsg(chat)
    }
  }
  return (
    <div className='flex items-center justify-center'>
      <input type="text" placeholder='Type message ...Hit enter to send' className='w-96 rounded-md p-2 text-blue-500' onChange={e =>(setChat(e.target.value))} onKeyDown={sendChat}/>
    </div>
  )
}

export default ChatInput