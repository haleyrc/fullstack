import React, { useState, useEffect }from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
	  <Pinger />
    </div>
  );
}

function Pinger() {
	const {status, error} = usePing()
	if (status === 'idle') {
		return <div>Ready</div>
	}
	if (status === 'fetching') {
		return <div>Pinging...</div>
	}
	if (status === 'error') {
		return <div>{error}</div>
	}
	return <div>Pong.</div>
}

function usePing() {
	const [status, setStatus] = useState('idle')
	const [error, setError] = useState(null)

	useEffect(() => {
		setStatus('fetching')
		fetch('/api/ping').then(resp => {
			console.group('Got response:')
			console.log(resp)
			console.groupEnd()
			setStatus('success')
		}).catch(e => {
			console.error(e)
			setError(e.message)
		})
	}, [])

	return {status, error}
}

export default App;
