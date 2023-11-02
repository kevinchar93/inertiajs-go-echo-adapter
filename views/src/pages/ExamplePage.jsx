import { Link } from "@inertiajs/react";

function ExamplePage({ phrase }) {

  return (
    <>
      <h1>{phrase}</h1>
      <Link href="/">Link to home page</Link>
    </>
  )
}

export default ExamplePage
