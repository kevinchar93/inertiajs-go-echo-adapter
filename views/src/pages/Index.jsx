import { useState } from 'react'
import reactLogo from '../assets/react.svg'
import viteLogo from '/vite.svg'
import './Index.css'
import { Link, Head } from "@inertiajs/react";

function Index({ header, allPokemon }) {
  const { title, meta } = header

  return (
    <div>
      <Head>
        <title>{title}</title>
        <meta name={meta.name} content={meta.content} />
      </Head>
      <h1>Simple Pokedex</h1>
      <ul>
        {allPokemon.map(({ id, name }) =>
          (<li key={id}>{id} - {name}</li>)
        )}
      </ul>
    </div>
  )
}

export default Index
