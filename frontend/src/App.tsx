// import { Button } from "@/components/ui/button"
// import React from "react"

function App() {
  const topBar = (
    <div className="h-10 bg-blue-500">Up bar</div>
  )

  const statusBar = (
    <div className="h-10 bg-red-500">Status bar</div>
  )

  const content = (
    <div className="flex-grow">Main part</div>
  )

  return (
    <div className="h-screen flex flex-col">
      {topBar}
      {content}
      {statusBar}
    </div>
  )
}

export default App
