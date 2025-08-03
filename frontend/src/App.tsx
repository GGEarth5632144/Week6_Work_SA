// App.tsx
import React from 'react';
import './App.css';

const App: React.FC = () => {
  return (
    <div className="container">
      {/* กล่องสีเขียวที่จัดตำแหน่งอยู่ตรงกลาง */}
      <div className="resizable-box">
        <p style={{ textAlign: 'center', marginTop: '0px' }}></p>
      </div>
    </div>
  );
}

export default App;
