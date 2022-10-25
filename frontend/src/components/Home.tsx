import { width } from '@mui/system';
import React from 'react'
import { Slide } from 'react-slideshow-image';
import 'react-slideshow-image/dist/styles.css'
const slideImages = [
  {
    url: 'images/libraryImage1.jpg',
    caption: 'Slide 1',
    width: 500,

  },
  {
    url: 'images/libraryImage2.jpg',
    caption: 'Slide 2',
    width: 500,
  },
];

function Home() {
  return (
    
          <div className="slide-container">
            <Slide>
             {slideImages.map((slideImage, index)=> (
                <div className="each-slide" key={index}>
                  <div style={{'backgroundImage': `url(${slideImage.url})`}}>
                    <span>{slideImage.caption}</span>
                  </div>
                </div>
              ))} 
            </Slide>
          </div>

  )
}

export default Home