import './Index.css'
import './Pokemon.css'
import { Link, Head } from "@inertiajs/react";

function Pokemon({ header, number, name, nextPokemon, prevPokemon, type, stats }) {
  const { title, meta } = header

  const paddedNumber = String(parseInt(number)).padStart(3, '0')

  return (
    <div style={{ width: '100%' }}>
      <div style={{ maxWidth: '600px', margin: 'auto' }}>
        <Head>
          <title>{title}</title>
          <meta name={meta.name} content={meta.content} />
        </Head>
        <h1>{name} (#{paddedNumber})</h1>
        <div className='img-container'>
          <img src={`/${paddedNumber}.png`} alt="" width={'250px'} height={'250px'} />
        </div>
        <h2>Type</h2>
        <ul>
          {type.map((t) =>
            (<li key={t}>{t}</li>)
          )}
        </ul>
        <h2>Stats</h2>
        <ul className='stats-list'>
          <li>HP: <p>{stats.hp}</p></li>
          <li>Attack: <p>{stats.attack}</p></li>
          <li>Defense: <p>{stats.defense}</p></li>
          <li>Special Attack: <p>{stats.spAttack}</p></li>
          <li>Special Defense: <p>{stats.spDefense}</p></li>
          <li>Speed: <p>{stats.speed}</p></li>
        </ul>
        <div className='links-container'>
          <span className='links'>
            {prevPokemon == -1 ? null : <Link href={`/pokemon/${prevPokemon}`}>Prev Pokemon</Link>}
            <Link href='/'>Back to Pokedex</Link>
            {nextPokemon == -1 ? null : <Link href={`/pokemon/${nextPokemon}`}>Next Pokemon</Link>}
          </span>
        </div>
      </div>
    </div>
  )
}

export default Pokemon
