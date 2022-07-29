## Soru 4
Aşağıdaki kod bloğu React uygulamalarında kullanılan `Routing` için App.js dosyası içinde oluşturulan  bir örnektir. Bu kod bloğuna baktığınızda:
* `exact` taglarını kullanmamış olsaydık ne olurdu?
* "/test2/" yolunu "/test1/"  için alt yol(sub route)yapmak istiyoruz. Bunu nasıl belirtiriz?
* Bu componentleri (Test, Test1, Test2, Test3) lazy load mantığıyla nasıl çağırabiliriz?_(Bonus Soru)_
> [React Router linkini](https://reactrouter.com/) kullanarak detaylı bilgilere ulaşabilirsiniz.
<br><br>
```javascript
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Test, Test1, Test2, Test3 } from './components';

import NavBarComponent from './components/NavBarComponent';

function App() {
    const baseUrl = document.getElementsByTagName('base')[0]?.getAttribute('href');

    return (
        <Router basename={baseUrl}>
            <NavBarComponent />
            <div>
                <Routes>
                    <Route exact path='/test/' element={<Test />} />
                    <Route exact path='/test1/' element={<Test1 />} />
                    <Route exact path='/test2/' element={<Test2 />} />
                    <Route exact path='/test3/' element={<Test3 />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
```
