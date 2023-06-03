
--seeds
INSERT INTO
    areas (name)
VALUES
    ('Al Sheikh zayed'),
    ('Al mohandseen');

INSERT INTO
    car_brands (name)
VALUES
    ('Toyota'),
    ('Mercides'),
    ('Hyundai');

INSERT INTO
    car_brand_models (name, years, car_brand_id)
VALUES
    (
        'Corolla',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        1
    ),
    (
        'Camry',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        1
    ),
    (
        'C200',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        2
    ),
    (
        'E200',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        2
    ),
    (
        'Elentra',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        3
    ),
    (
        'Tucson',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019',
        3
    ),
    (
        'ix35',
        '2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022',
        3
    );

INSERT INTO
    centers (
        name,
        phone,
        address,
        area_id,
        lat,
        long
    )
VALUES
    (
        'GB Auto - Hyundai - El Sheikh Zayed',
        '01207661993',
        'خلف الهاييبر وان بجوار امريكانا بلازا - الحى الاول شارع الروضة امام كليه, الهندسة, Giza, الجيزة',
        1,
        30.03735564,
        31.01083698
    ),
    (
        'مركز صيانة هيونداى Hyundai Service Center',
        '01006888826',
        'بوابه, 8 بيفرلى هيلز، الجيزة 12588',
        1,
        30.06500699,
        30.9572968
    ),
    (
        'GB Auto - Hyundai - Abo Rawash - Showroom, Service Center & spare parts',
        '0235366010',
        'ابورواش 28 طريق القاهرة الاسكندرية الصحراوى - المنطقة الصناعية الثانية - طريق ال, Abu Rawash, الجيزة',
        1,
        30.067845596254937,
        31.03740789
    ),
    (
        'Wahdan Auto Group',
        '0238204410',
        'المنطقة الصناعية الثالثة، السادس من اكتوبر، الجيزة، 3221732',
        1,
        29.914065410175006,
        30.8949261
    ),
    (
        'مركز عربيتي',
        '01157555959',
        'رقم 45، قسم الشيخ زايد، الجيزة 3242803',
        1,
        30.043544240746243,
        30.96450447
    );

INSERT INTO
    "winch" (
        "area_id",
        "name",
        "phone",
        "driver_name",
        "lat",
        "long",
        "driver_phone"
    )
VALUES
    (
        1,
        'Winch 1',
        '1234567890',
        'Driver 1',
        30.043,
        31.121,
        '1234567890'
    ),
    (
        1,
        'Winch 2',
        '1234567890',
        'Driver 2',
        31.043,
        32.121,
        '1234567890'
    ),
    (
        1,
        'Winch 3',
        '1234567890',
        'Driver 3',
        32.043,
        33.121,
        '1234567890'
    ),
    (
        2,
        'Winch 4',
        '1234567890',
        'Driver 4',
        33.043,
        34.121,
        '1234567890'
    ),
    (
        2,
        'Winch 5',
        '1234567890',
        'Driver 5',
        34.043,
        35.121,
        '1234567890'
    );

INSERT INTO codes(
    car_brand_model_id,
    vehicle_part,
    code_name,
    description,
    is_emergency,
    code_type
) VALUES
    (
        1,
        'Engine (engine control)',
        'P0118',
        'Engine Coolant Temperature Circuit High Input',
        true,
        'M'
    )
    ,
 
    (
        1,
        'Engine (engine control)',
        'P0132',
        'O2 Sensor Circuit High Voltage (Bank 1 / Sensor 1)',
        false,
        'M'
    ),
    (
        1,
        'Engine (engine control)',
        'P0221',
        'Throttle/Pedal Position Sensor/Switch \"B\" Circuit Range/Performance',
        false,
        'M'
    ),
    (
        1,
        'Engine (engine control)',
        'P0230',
        'Fuel Pump Primary Circuit',
        false,
        'M'
    ),
    (
        1,
        'Engine (engine control)',
        'P0301',
        'Cylinder 1-Misfire detected',
        false,
        'M'
    ),
    (
        1,
        'Engine (engine control)',
        'P250F',
        'Engine Oil Level Too Low',
        true,
        'M'
    ),
    (
        1,
        'Transmission',
        'P0712',
        'Transmission Fluid Temperature Sensor \"A\" Circuit Low Input',
        true,
        'M'
    ),
    (
        1,
        'Transmission',
        'P0731',
        'Gear 1 Incorrect Ratio',
        false,
        'M'
    ),
    (
        1,
        'Transmission',
        'P0753',
        'Shift Control Solenoid Valve \"A\" Electrical (UD/B)',
        false,
        'M'
    ),
    (
        1,
        'Transmission',
        'P0955',
        'Auto Shift Manual Mode Circuit',
        false,
        'M'
    ),
    (
        1,
        'Air Conditioner',
        'B1246',
        'Air Mix Potentiometer Short (High)-Driver',
        false,
        'M'
    ),
    (
        1,
        'Air Conditioner',
        'B1237',
        'Ambient Temperature Sensor Short (Low)',
        false,
        'B'
    ),
    (
        1,
        'Airbag',
        'B1346',
        'Driver Airbag Resistance Too High (1st stage)',
        false,
        'B'
    ),
    (
        1,
        'Airbag',
        'B1361',
        'Pretensioner Front-Driver Resistance too High',
        false,
        'B'
    ),
    (
        1,
        'Brake /ABS/ESP',
        'C1200',
        'Wheel Speed Sensor Front-LH Open/Short',
        false,
        'M'
    ),
    (
        1,
        'Brake /ABS/ES',
        'C1503',
        'TCS/ESC(ESP) Switch Error',
        false,
        'M'
    ),
    (
        1,
        'Brake /ABS/ES',
        'C2380',
        'ABS/TCS/ESC(ESP) Valve Error',
        false,
        'M'
    ),
    (
        1,
        'Brake /ABS/ES',
        'C1616',
        'CAN Bus OFF',
        false,
        'M'
    ),
   
    (
        1,
        'Steering',
        'C1259',
        'Steering Angle Sensor – Electrical Malfunction',
        false,
        'M'
    ),
    (
        1,
        'Steering',
        'C1261',
        'Steering Angle Sensor Not Calibrated',
        false,
        'M'
    ),
    (
        1,
        'Tire Pressure Monitoring',
        'C1212',
        'Vehicle Speed Sensor',
        false,
        'M'
    ),
    (
        1,
        'Tire Pressure Monitoring',
        'C1660',
        'Receiver RF Circuit Fail',
        false,
        'T'
    ),
    (
        1,
        'Parking System',
        'C1368',
        'PAS Sensor - Rear Left Outer',
        false,
        'B'
    ),
    (
        1,
        'Parking System',
        'C2232',
        'SPAS Speaker Failure',
        false,
        'B'
    ),
     (
        2,
        'Engine (engine control)',
        'P0118',
        'Engine Coolant Temperature Circuit High Input',
        true,
        'M'
    )
    ,
 
    (
        2,
        'Engine (engine control)',
        'P0132',
        'O2 Sensor Circuit High Voltage (Bank 1 / Sensor 1)',
        false,
        'M'
    ),
    (
        2,
        'Engine (engine control)',
        'P0221',
        'Throttle/Pedal Position Sensor/Switch \"B\" Circuit Range/Performance',
        false,
        'M'
    ),
    (
        2,
        'Engine (engine control)',
        'P0230',
        'Fuel Pump Primary Circuit',
        false,
        'M'
    ),
    (
        2,
        'Engine (engine control)',
        'P0301',
        'Cylinder 1-Misfire detected',
        false,
        'M'
    ),
    (
        2,
        'Engine (engine control)',
        'P250F',
        'Engine Oil Level Too Low',
        true,
        'M'
    ),
    (
        2,
        'Transmission',
        'P0712',
        'Transmission Fluid Temperature Sensor \"A\" Circuit Low Input',
        true,
        'M'
    ),
    (
        2,
        'Transmission',
        'P0731',
        'Gear 1 Incorrect Ratio',
        false,
        'M'
    ),
    (
        2,
        'Transmission',
        'P0753',
        'Shift Control Solenoid Valve \"A\" Electrical (UD/B)',
        false,
        'M'
    ),
    (
        2,
        'Transmission',
        'P0955',
        'Auto Shift Manual Mode Circuit',
        false,
        'M'
    ),
    (
        2,
        'Air Conditioner',
        'B1246',
        'Air Mix Potentiometer Short (High)-Driver',
        false,
        'M'
    ),
    (
        2,
        'Air Conditioner',
        'B1237',
        'Ambient Temperature Sensor Short (Low)',
        false,
        'B'
    ),
    (
        2,
        'Airbag',
        'B1346',
        'Driver Airbag Resistance Too High (1st stage)',
        false,
        'B'
    ),
    (
        2,
        'Airbag',
        'B1361',
        'Pretensioner Front-Driver Resistance too High',
        false,
        'B'
    ),
    (
        2,
        'Brake /ABS/ESP',
        'C1200',
        'Wheel Speed Sensor Front-LH Open/Short',
        false,
        'M'
    ),
    (
        2,
        'Brake /ABS/ES',
        'C1503',
        'TCS/ESC(ESP) Switch Error',
        false,
        'M'
    ),
    (
        2,
        'Brake /ABS/ES',
        'C2380',
        'ABS/TCS/ESC(ESP) Valve Error',
        false,
        'M'
    ),
    (
        2,
        'Brake /ABS/ES',
        'C1616',
        'CAN Bus OFF',
        false,
        'M'
    ),
   
    (
        2,
        'Steering',
        'C1259',
        'Steering Angle Sensor – Electrical Malfunction',
        false,
        'M'
    ),
    (
        2,
        'Steering',
        'C1261',
        'Steering Angle Sensor Not Calibrated',
        false,
        'M'
    ),
    (
        2,
        'Tire Pressure Monitoring',
        'C1212',
        'Vehicle Speed Sensor',
        false,
        'M'
    ),
    (
        2,
        'Tire Pressure Monitoring',
        'C1660',
        'Receiver RF Circuit Fail',
        false,
        'T'
    ),
    (
        2,
        'Parking System',
        'C1368',
        'PAS Sensor - Rear Left Outer',
        false,
        'B'
    ),
    (
        2,
        'Parking System',
        'C2232',
        'SPAS Speaker Failure',
        false,
        'B'
    ),
     (
       3,
        'Engine (engine control)',
        'P0118',
        'Engine Coolant Temperature Circuit High Input',
        true,
        'M'
    )
    ,
 
    (
       3,
        'Engine (engine control)',
        'P0132',
        'O2 Sensor Circuit High Voltage (Bank 1 / Sensor 1)',
        false,
        'M'
    ),
    (
       3,
        'Engine (engine control)',
        'P0221',
        'Throttle/Pedal Position Sensor/Switch \"B\" Circuit Range/Performance',
        false,
        'M'
    ),
    (
       3,
        'Engine (engine control)',
        'P0230',
        'Fuel Pump Primary Circuit',
        false,
        'M'
    ),
    (
       3,
        'Engine (engine control)',
        'P0301',
        'Cylinder 1-Misfire detected',
        false,
        'M'
    ),
    (
       3,
        'Engine (engine control)',
        'P250F',
        'Engine Oil Level Too Low',
        true,
        'M'
    ),
    (
       3,
        'Transmission',
        'P0712',
        'Transmission Fluid Temperature Sensor \"A\" Circuit Low Input',
        true,
        'M'
    ),
    (
       3,
        'Transmission',
        'P0731',
        'Gear 1 Incorrect Ratio',
        false,
        'M'
    ),
    (
       3,
        'Transmission',
        'P0753',
        'Shift Control Solenoid Valve \"A\" Electrical (UD/B)',
        false,
        'M'
    ),
    (
       3,
        'Transmission',
        'P0955',
        'Auto Shift Manual Mode Circuit',
        false,
        'M'
    ),
    (
       3,
        'Air Conditioner',
        'B1246',
        'Air Mix Potentiometer Short (High)-Driver',
        false,
        'M'
    ),
    (
       3,
        'Air Conditioner',
        'B1237',
        'Ambient Temperature Sensor Short (Low)',
        false,
        'B'
    ),
    (
       3,
        'Airbag',
        'B1346',
        'Driver Airbag Resistance Too High (1st stage)',
        false,
        'B'
    ),
    (
       3,
        'Airbag',
        'B1361',
        'Pretensioner Front-Driver Resistance too High',
        false,
        'B'
    ),
    (
       3,
        'Brake /ABS/ESP',
        'C1200',
        'Wheel Speed Sensor Front-LH Open/Short',
        false,
        'M'
    ),
    (
       3,
        'Brake /ABS/ES',
        'C1503',
        'TCS/ESC(ESP) Switch Error',
        false,
        'M'
    ),
    (
       3,
        'Brake /ABS/ES',
        'C2380',
        'ABS/TCS/ESC(ESP) Valve Error',
        false,
        'M'
    ),
    (
       3,
        'Brake /ABS/ES',
        'C1616',
        'CAN Bus OFF',
        false,
        'M'
    ),
   
    (
       3,
        'Steering',
        'C1259',
        'Steering Angle Sensor – Electrical Malfunction',
        false,
        'M'
    ),
    (
       3,
        'Steering',
        'C1261',
        'Steering Angle Sensor Not Calibrated',
        false,
        'M'
    ),
    (
       3,
        'Tire Pressure Monitoring',
        'C1212',
        'Vehicle Speed Sensor',
        false,
        'M'
    ),
    (
       3,
        'Tire Pressure Monitoring',
        'C1660',
        'Receiver RF Circuit Fail',
        false,
        'T'
    ),
    (
       3,
        'Parking System',
        'C1368',
        'PAS Sensor - Rear Left Outer',
        false,
        'B'
    ),
    (
       3,
        'Parking System',
        'C2232',
        'SPAS Speaker Failure',
        false,
        'B'
    ),
     (
       4,
        'Engine (engine control)',
        'P0118',
        'Engine Coolant Temperature Circuit High Input',
        true,
        'M'
    )
    ,
 
    (
       4,
        'Engine (engine control)',
        'P0132',
        'O2 Sensor Circuit High Voltage (Bank 1 / Sensor 1)',
        false,
        'M'
    ),
    (
       4,
        'Engine (engine control)',
        'P0221',
        'Throttle/Pedal Position Sensor/Switch \"B\" Circuit Range/Performance',
        false,
        'M'
    ),
    (
       4,
        'Engine (engine control)',
        'P0230',
        'Fuel Pump Primary Circuit',
        false,
        'M'
    ),
    (
       4,
        'Engine (engine control)',
        'P0301',
        'Cylinder 1-Misfire detected',
        false,
        'M'
    ),
    (
       4,
        'Engine (engine control)',
        'P250F',
        'Engine Oil Level Too Low',
        true,
        'M'
    ),
    (
       4,
        'Transmission',
        'P0712',
        'Transmission Fluid Temperature Sensor \"A\" Circuit Low Input',
        true,
        'M'
    ),
    (
       4,
        'Transmission',
        'P0731',
        'Gear 1 Incorrect Ratio',
        false,
        'M'
    ),
    (
       4,
        'Transmission',
        'P0753',
        'Shift Control Solenoid Valve \"A\" Electrical (UD/B)',
        false,
        'M'
    ),
    (
       4,
        'Transmission',
        'P0955',
        'Auto Shift Manual Mode Circuit',
        false,
        'M'
    ),
    (
       4,
        'Air Conditioner',
        'B1246',
        'Air Mix Potentiometer Short (High)-Driver',
        false,
        'M'
    ),
    (
       4,
        'Air Conditioner',
        'B1237',
        'Ambient Temperature Sensor Short (Low)',
        false,
        'B'
    ),
    (
       4,
        'Airbag',
        'B1346',
        'Driver Airbag Resistance Too High (1st stage)',
        false,
        'B'
    ),
    (
       4,
        'Airbag',
        'B1361',
        'Pretensioner Front-Driver Resistance too High',
        false,
        'B'
    ),
    (
       4,
        'Brake /ABS/ESP',
        'C1200',
        'Wheel Speed Sensor Front-LH Open/Short',
        false,
        'M'
    ),
    (
       4,
        'Brake /ABS/ES',
        'C1503',
        'TCS/ESC(ESP) Switch Error',
        false,
        'M'
    ),
    (
       4,
        'Brake /ABS/ES',
        'C2380',
        'ABS/TCS/ESC(ESP) Valve Error',
        false,
        'M'
    ),
    (
       4,
        'Brake /ABS/ES',
        'C1616',
        'CAN Bus OFF',
        false,
        'M'
    ),
   
    (
       4,
        'Steering',
        'C1259',
        'Steering Angle Sensor – Electrical Malfunction',
        false,
        'M'
    ),
    (
       4,
        'Steering',
        'C1261',
        'Steering Angle Sensor Not Calibrated',
        false,
        'M'
    ),
    (
       4,
        'Tire Pressure Monitoring',
        'C1212',
        'Vehicle Speed Sensor',
        false,
        'M'
    ),
    (
       4,
        'Tire Pressure Monitoring',
        'C1660',
        'Receiver RF Circuit Fail',
        false,
        'T'
    ),
    (
       4,
        'Parking System',
        'C1368',
        'PAS Sensor - Rear Left Outer',
        false,
        'B'
    ),
    (
       4,
        'Parking System',
        'C2232',
        'SPAS Speaker Failure',
        false,
        'B'
    ),
    (
       5,
        'Engine (engine control)',
        'P0118',
        'Engine Coolant Temperature Circuit High Input',
        true,
        'M'
    )
    ,
 
    (
       5,
        'Engine (engine control)',
        'P0132',
        'O2 Sensor Circuit High Voltage (Bank 1 / Sensor 1)',
        false,
        'M'
    ),
    (
       5,
        'Engine (engine control)',
        'P0221',
        'Throttle/Pedal Position Sensor/Switch \"B\" Circuit Range/Performance',
        false,
        'M'
    ),
    (
       5,
        'Engine (engine control)',
        'P0230',
        'Fuel Pump Primary Circuit',
        false,
        'M'
    ),
    (
       5,
        'Engine (engine control)',
        'P0301',
        'Cylinder 1-Misfire detected',
        false,
        'M'
    ),
    (
       5,
        'Engine (engine control)',
        'P250F',
        'Engine Oil Level Too Low',
        true,
        'M'
    ),
    (
       5,
        'Transmission',
        'P0712',
        'Transmission Fluid Temperature Sensor \"A\" Circuit Low Input',
        true,
        'M'
    ),
    (
       5,
        'Transmission',
        'P0731',
        'Gear 1 Incorrect Ratio',
        false,
        'M'
    ),
    (
       5,
        'Transmission',
        'P0753',
        'Shift Control Solenoid Valve \"A\" Electrical (UD/B)',
        false,
        'M'
    ),
    (
       5,
        'Transmission',
        'P0955',
        'Auto Shift Manual Mode Circuit',
        false,
        'M'
    ),
    (
       5,
        'Air Conditioner',
        'B1246',
        'Air Mix Potentiometer Short (High)-Driver',
        false,
        'M'
    ),
    (
       5,
        'Air Conditioner',
        'B1237',
        'Ambient Temperature Sensor Short (Low)',
        false,
        'B'
    ),
    (
       5,
        'Airbag',
        'B1346',
        'Driver Airbag Resistance Too High (1st stage)',
        false,
        'B'
    ),
    (
       5,
        'Airbag',
        'B1361',
        'Pretensioner Front-Driver Resistance too High',
        false,
        'B'
    ),
    (
       5,
        'Brake /ABS/ESP',
        'C1200',
        'Wheel Speed Sensor Front-LH Open/Short',
        false,
        'M'
    ),
    (
       5,
        'Brake /ABS/ES',
        'C1503',
        'TCS/ESC(ESP) Switch Error',
        false,
        'M'
    ),
    (
       5,
        'Brake /ABS/ES',
        'C2380',
        'ABS/TCS/ESC(ESP) Valve Error',
        false,
        'M'
    ),
    (
       5,
        'Brake /ABS/ES',
        'C1616',
        'CAN Bus OFF',
        false,
        'M'
    ),
   
    (
       5,
        'Steering',
        'C1259',
        'Steering Angle Sensor – Electrical Malfunction',
        false,
        'M'
    ),
    (
       5,
        'Steering',
        'C1261',
        'Steering Angle Sensor Not Calibrated',
        false,
        'M'
    ),
    (
       5,
        'Tire Pressure Monitoring',
        'C1212',
        'Vehicle Speed Sensor',
        false,
        'M'
    ),
    (
       5,
        'Tire Pressure Monitoring',
        'C1660',
        'Receiver RF Circuit Fail',
        false,
        'T'
    ),
    (
       5,
        'Parking System',
        'C1368',
        'PAS Sensor - Rear Left Outer',
        false,
        'B'
    ),
    (
       5,
        'Parking System',
        'C2232',
        'SPAS Speaker Failure',
        false,
        'B'
    ),
    (
       7,
        'Parking System',
        'P0420',
        'low catalyst system efficiency',
        false,
        'B'
    ),
    (
       7,
        'Parking System',
        'P0421',
        'relates to the catalytic converter in your  exhaust system',
        false,
        'B'
    );


-- --insert default user
--   SELECT  *  FROM user_create('ahmed','+20107110101','ahmed@obd.com','$2a$10$Mot8FRSl2y5yBu2fdg.C3eLOdBXd2dZz/v/V004Rlia1bztMRyU4e',7,2022);


