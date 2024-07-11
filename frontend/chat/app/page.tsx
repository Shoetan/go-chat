"use client"

import { connect, sendMsg } from "@/api/api";
import ChatHistory from "@/components/ChatHistory";
import ChatInput from "@/components/ChatInput";
import Header from "@/components/Header";
import Image from "next/image";
import { useEffect, useState } from "react";

export default function Home() {

  const message = [1,2,3,4]

  const [chatHistory, setChatHistory] = useState<any>([])

  useEffect(()=>{
    const handleNewMessage = (msg:any) =>{
      console.log("New Message")
      setChatHistory((prev: any) => [...prev, msg])
    }
    connect(handleNewMessage)
  },[])

  const text =  chatHistory?.map((event: { data: any; }) => event.data)

  const bodyArray = text.map((jsonString: string) => JSON.parse(jsonString).body);

  return (
    <main className="">
     <Header/>
     <ChatHistory messages={bodyArray}/>
     <ChatInput/>
    </main>
  );
}
