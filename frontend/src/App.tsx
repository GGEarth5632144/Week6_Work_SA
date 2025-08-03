import React, { useState } from 'react';
import './App.css';
import image1 from './assets/images/image-5-6.png';
import image2 from './assets/images/image-2-2.png';

import {
  CaretRightFilled,
} from '@ant-design/icons';
import {  Menu } from 'antd';
import { type GetProp, type MenuProps, Col,  Row } from 'antd';

type MenuItem = GetProp<MenuProps, 'items'>[number];

const items: MenuItem[] = [
  {
    key: '1',
    icon: <CaretRightFilled />,
    label: 'ลงทะเบียนเรียน',
  },
  {
    key: '2',
    icon: <CaretRightFilled />,
    label: 'วิชาที่เปิดสอน',
  },
  {
    key: '3',
    icon: <CaretRightFilled />,
    label: 'ตารางเรียน',
  },
  {
    key: '4',
    icon: <CaretRightFilled />,
    label: 'ผลการเรียน',
  },
  {
    key: '5',
    icon: <CaretRightFilled />,
    label: 'คะแนน',
  },
  {
    key: '6',
    icon: <CaretRightFilled />,
    label: 'ใบแจ้งยอดชำระ',
  },
  {
    key: '7',
    icon: <CaretRightFilled />,
    label: 'ระเบียนประวัติ',
  },
  {
    key: '8',
    icon: <CaretRightFilled />,
    label: 'อาจารย์',
  },
  {
    key: '9',
    icon: <CaretRightFilled />,
    label: 'คำร้อง',
  },
  {
    key: '10',
    icon: <CaretRightFilled />,
    label: 'แจ้งจบการศึกษา',
  },
  {
    key: '11',
    icon: <CaretRightFilled />,
    label: 'หลักสูตร',
  },
  {
    key: '12',
    icon: <CaretRightFilled />,
    label: 'เปลี่ยนรหัสผ่าน',
  },
];

const App: React.FC = () => {
  return (
    <div className="app-wrapper">
      <div className="container1">
        <div className="resizable-box1">
          <div className="containernavigationtab">
            <Menu
              className="custom-menu"
              style={{
                      backgroundColor: '#bda0a000', // สีพื้นหลังเมนู
                      borderRadius: '8px',
                      width: 270,
                      height: 100,
                    }}
              defaultSelectedKeys={['1']}
              defaultOpenKeys={['sub1']}
              mode="vertical" // คงที่ ไม่ให้ผู้ใช้เลือก
              items={items}
            />
          </div>
        </div>
      </div>
      <Row style={{ backgroundColor: '#003eb3', padding: '10px' ,width: '1500px'}}>
        <Col span={8} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
          <h1 style={{ color: 'white', fontSize: '24px', margin: 0 }}>ยินดีต้อนรับ คุณ ณัฐนันท์ วงษ์หาญ</h1>
        </Col>
        <Col span={8}></Col>
        <Col span={7}>
          <img src={image1} alt="โลโก้ชื่อมหาลัย" style={{ width: '200px', height: 'auto' ,float: 'right' }} />
        </Col>
      </Row>
      <div className='container2' >
        <img src={image2} alt="โลโก้มหาลัย" style={{ width: '300px', height: 'auto' ,float: 'right' }} />
      </div>
    </div>
  );
};


export default App;
