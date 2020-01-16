import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';

import Routes from './Routes';

const App = () => {
    return(
        <BrowserRouter>
            <Routes/>
        </BrowserRouter>
    );
};

// load App whenever root HTML element is requested
ReactDOM.render(<App/>, document.getElementById('root'));
