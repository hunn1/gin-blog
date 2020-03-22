// Utility function
function Util () {};

/* 
	class manipulation functions
*/
Util.hasClass = function(el, className) {
	if (el.classList) return el.classList.contains(className);
	else return !!el.className.match(new RegExp('(\\s|^)' + className + '(\\s|$)'));
};

Util.addClass = function(el, className) {
	var classList = className.split(' ');
 	if (el.classList) el.classList.add(classList[0]);
 	else if (!Util.hasClass(el, classList[0])) el.className += " " + classList[0];
 	if (classList.length > 1) Util.addClass(el, classList.slice(1).join(' '));
};

Util.removeClass = function(el, className) {
	var classList = className.split(' ');
	if (el.classList) el.classList.remove(classList[0]);	
	else if(Util.hasClass(el, classList[0])) {
		var reg = new RegExp('(\\s|^)' + classList[0] + '(\\s|$)');
		el.className=el.className.replace(reg, ' ');
	}
	if (classList.length > 1) Util.removeClass(el, classList.slice(1).join(' '));
};

Util.toggleClass = function(el, className, bool) {
	if(bool) Util.addClass(el, className);
	else Util.removeClass(el, className);
};

Util.setAttributes = function(el, attrs) {
  for(var key in attrs) {
    el.setAttribute(key, attrs[key]);
  }
};

/* 
  DOM manipulation
*/
Util.getChildrenByClassName = function(el, className) {
  var children = el.children,
    childrenByClass = [];
  for (var i = 0; i < el.children.length; i++) {
    if (Util.hasClass(el.children[i], className)) childrenByClass.push(el.children[i]);
  }
  return childrenByClass;
};

Util.is = function(elem, selector) {
  if(selector.nodeType){
    return elem === selector;
  }

  var qa = (typeof(selector) === 'string' ? document.querySelectorAll(selector) : selector),
    length = qa.length,
    returnArr = [];

  while(length--){
    if(qa[length] === elem){
      return true;
    }
  }

  return false;
};

/* 
	Animate height of an element
*/
Util.setHeight = function(start, to, element, duration, cb) {
	var change = to - start,
	    currentTime = null;

  var animateHeight = function(timestamp){  
    if (!currentTime) currentTime = timestamp;         
    var progress = timestamp - currentTime;
    var val = parseInt((progress/duration)*change + start);
    element.style.height = val+"px";
    if(progress < duration) {
        window.requestAnimationFrame(animateHeight);
    } else {
    	cb();
    }
  };
  
  //set the height of the element before starting animation -> fix bug on Safari
  element.style.height = start+"px";
  window.requestAnimationFrame(animateHeight);
};

/* 
	Smooth Scroll
*/

Util.scrollTo = function(final, duration, cb) {
  var start = window.scrollY || document.documentElement.scrollTop,
      currentTime = null;
      
  var animateScroll = function(timestamp){
  	if (!currentTime) currentTime = timestamp;        
    var progress = timestamp - currentTime;
    if(progress > duration) progress = duration;
    var val = Math.easeInOutQuad(progress, start, final-start, duration);
    window.scrollTo(0, val);
    if(progress < duration) {
        window.requestAnimationFrame(animateScroll);
    } else {
      cb && cb();
    }
  };

  window.requestAnimationFrame(animateScroll);
};

/* 
  Focus utility classes
*/

//Move focus to an element
Util.moveFocus = function (element) {
  if( !element ) element = document.getElementsByTagName("body")[0];
  element.focus();
  if (document.activeElement !== element) {
    element.setAttribute('tabindex','-1');
    element.focus();
  }
};

/* 
  Misc
*/

Util.getIndexInArray = function(array, el) {
  return Array.prototype.indexOf.call(array, el);
};

Util.cssSupports = function(property, value) {
  if('CSS' in window) {
    return CSS.supports(property, value);
  } else {
    var jsProperty = property.replace(/-([a-z])/g, function (g) { return g[1].toUpperCase();});
    return jsProperty in document.body.style;
  }
};

// merge a set of user options into plugin defaults
// https://gomakethings.com/vanilla-javascript-version-of-jquery-extend/
Util.extend = function() {
  // Variables
  var extended = {};
  var deep = false;
  var i = 0;
  var length = arguments.length;

  // Check if a deep merge
  if ( Object.prototype.toString.call( arguments[0] ) === '[object Boolean]' ) {
    deep = arguments[0];
    i++;
  }

  // Merge the object into the extended object
  var merge = function (obj) {
    for ( var prop in obj ) {
      if ( Object.prototype.hasOwnProperty.call( obj, prop ) ) {
        // If deep merge and property is an object, merge properties
        if ( deep && Object.prototype.toString.call(obj[prop]) === '[object Object]' ) {
          extended[prop] = extend( true, extended[prop], obj[prop] );
        } else {
          extended[prop] = obj[prop];
        }
      }
    }
  };

  // Loop through each object and conduct a merge
  for ( ; i < length; i++ ) {
    var obj = arguments[i];
    merge(obj);
  }

  return extended;
};

// Check if Reduced Motion is enabled
Util.osHasReducedMotion = function() {
  var matchMediaObj = window.matchMedia('(prefers-reduced-motion: reduce)');
  if(matchMediaObj) return matchMediaObj.matches;
  return false; // return false if not supported
}; 

/* 
	Polyfills
*/
//Closest() method
if (!Element.prototype.matches) {
	Element.prototype.matches = Element.prototype.msMatchesSelector || Element.prototype.webkitMatchesSelector;
}

if (!Element.prototype.closest) {
	Element.prototype.closest = function(s) {
		var el = this;
		if (!document.documentElement.contains(el)) return null;
		do {
			if (el.matches(s)) return el;
			el = el.parentElement || el.parentNode;
		} while (el !== null && el.nodeType === 1); 
		return null;
	};
}

//Custom Event() constructor
if ( typeof window.CustomEvent !== "function" ) {

  function CustomEvent ( event, params ) {
    params = params || { bubbles: false, cancelable: false, detail: undefined };
    var evt = document.createEvent( 'CustomEvent' );
    evt.initCustomEvent( event, params.bubbles, params.cancelable, params.detail );
    return evt;
   }

  CustomEvent.prototype = window.Event.prototype;

  window.CustomEvent = CustomEvent;
}

/* 
	Animation curves
*/
Math.easeInOutQuad = function (t, b, c, d) {
	t /= d/2;
	if (t < 1) return c/2*t*t + b;
	t--;
	return -c/2 * (t*(t-2) - 1) + b;
};
// File#: _1_accordion
(function() {
	var Accordion = function(element) {
		this.element = element;
		this.items = Util.getChildrenByClassName(this.element, 'js-accordion__item');
		this.showClass = 'accordion__item--is-open';
		this.animateHeight = (this.element.getAttribute('data-animation') == 'on');
		this.multiItems = !(this.element.getAttribute('data-multi-items') == 'off'); 
		this.initAccordion();
	};

	Accordion.prototype.initAccordion = function() {
		//set initial aria attributes
		for( var i = 0; i < this.items.length; i++) {
			var button = this.items[i].getElementsByTagName('button')[0],
				content = this.items[i].getElementsByClassName('js-accordion__panel')[0],
				isOpen = Util.hasClass(this.items[i], this.showClass) ? 'true' : 'false';
			Util.setAttributes(button, {'aria-expanded': isOpen, 'aria-controls': 'accordion-content-'+i, 'id': 'accordion-header-'+i});
			Util.addClass(button, 'js-accordion__trigger');
			Util.setAttributes(content, {'aria-labelledby': 'accordion-header-'+i, 'id': 'accordion-content-'+i});
		}

		//listen for Accordion events
		this.initAccordionEvents();
	};

	Accordion.prototype.initAccordionEvents = function() {
		var self = this;

		this.element.addEventListener('click', function(event) {
			var trigger = event.target.closest('.js-accordion__trigger');
			//check index to make sure the click didn't happen inside a children accordion
			if( trigger && Util.getIndexInArray(self.items, trigger.parentElement) >= 0) self.triggerAccordion(trigger);
		});
	};

	Accordion.prototype.triggerAccordion = function(trigger) {
		var self = this;
		var bool = (trigger.getAttribute('aria-expanded') === 'true');

		this.animateAccordion(trigger, bool);
	};

	Accordion.prototype.animateAccordion = function(trigger, bool) {
		var self = this;
		var item = trigger.closest('.js-accordion__item'),
			content = item.getElementsByClassName('js-accordion__panel')[0],
			ariaValue = bool ? 'false' : 'true';

		if(!bool) Util.addClass(item, this.showClass);
		trigger.setAttribute('aria-expanded', ariaValue);

		if(this.animateHeight) {
			//store initial and final height - animate accordion content height
			var initHeight = bool ? content.offsetHeight: 0,
				finalHeight = bool ? 0 : content.offsetHeight;
		}

		if(window.requestAnimationFrame && this.animateHeight) {
			Util.setHeight(initHeight, finalHeight, content, 200, function(){
				self.resetContentVisibility(item, content, bool);
			});
		} else {
			self.resetContentVisibility(item, content, bool);
		}

		if( !this.multiItems && !bool) this.closeSiblings(item);

	};

	Accordion.prototype.resetContentVisibility = function(item, content, bool) {
		Util.toggleClass(item, this.showClass, !bool);
		content.removeAttribute("style");
	};

	Accordion.prototype.closeSiblings = function(item) {
		//if only one accordion can be open -> search if there's another one open
		var index = Util.getIndexInArray(this.items, item);
		for( var i = 0; i < this.items.length; i++) {
			if(Util.hasClass(this.items[i], this.showClass) && i != index) {
				this.animateAccordion(this.items[i].getElementsByClassName('js-accordion__trigger')[0], true);
				return false;
			}
		}
	};
	
	//initialize the Accordion objects
	var accordions = document.getElementsByClassName('js-accordion');
	if( accordions.length > 0 ) {
		for( var i = 0; i < accordions.length; i++) {
			(function(i){new Accordion(accordions[i]);})(i);
		}
	}
}());
// File#: _1_contact
/*
	⚠️ Make sure to include the Google Maps API. 
	You can include the script right after the contact.js (before the body closing tag in the index.html file):
	<script async defer src="https://maps.googleapis.com/maps/api/js?key=YOUR_API_KEY&callback=initGoogleMap"></script> 
*/
function initGoogleMap() {
	var contactMap = document.getElementsByClassName('js-contact__map');
	if(contactMap.length > 0) {
		for(var i = 0; i < contactMap.length; i++) {
			initContactMap(contactMap[i]);
		}
	}
};

function initContactMap(wrapper) {
	var coordinate = wrapper.getAttribute('data-coordinates').split(',');
	var map = new google.maps.Map(wrapper, {zoom: 10, center: {lat: Number(coordinate[0]), lng:  Number(coordinate[1])}});
	var marker = new google.maps.Marker({position: {lat: Number(coordinate[0]), lng:  Number(coordinate[1])}, map: map});
};
// File#: _1_main-header
(function() {
	var mainHeader = document.getElementsByClassName('js-main-header')[0];
	if( mainHeader ) {
		var trigger = mainHeader.getElementsByClassName('js-main-header__nav-trigger')[0],
			nav = mainHeader.getElementsByClassName('js-main-header__nav')[0];
		//detect click on nav trigger
		trigger.addEventListener("click", function(event) {
			event.preventDefault();
			var ariaExpanded = !Util.hasClass(nav, 'main-header__nav--is-visible');
			//show nav and update button aria value
			Util.toggleClass(nav, 'main-header__nav--is-visible', ariaExpanded);
			trigger.setAttribute('aria-expanded', ariaExpanded);
			if(ariaExpanded) { //opening menu -> move focus to first element inside nav
				nav.querySelectorAll('[href], input:not([disabled]), button:not([disabled])')[0].focus();
			}
		});
	}
}());
// File#: _1_reading-progressbar
(function() {
	var readingIndicator = document.getElementsByClassName('js-reading-progressbar')[0],
		readingIndicatorContent = document.getElementsByClassName('js-reading-content')[0];
	
	if( readingIndicator && readingIndicatorContent) {
		var progressInfo = [],
			progressEvent = false,
			progressFallback = readingIndicator.getElementsByClassName('js-reading-progressbar__fallback')[0],
			progressIsSupported = 'value' in readingIndicator;

		progressInfo['height'] = readingIndicatorContent.offsetHeight;
		progressInfo['top'] = readingIndicatorContent.getBoundingClientRect().top;
		progressInfo['window'] = window.innerHeight;
		progressInfo['class'] = 'reading-progressbar--is-active';
		
		//init indicator
		setProgressIndicator();
		//listen to the window scroll event - update progress
		window.addEventListener('scroll', function(event){
			if(progressEvent) return;
			progressEvent = true;
			(!window.requestAnimationFrame) ? setTimeout(function(){setProgressIndicator();}, 250) : window.requestAnimationFrame(setProgressIndicator);
		});
		// listen to window resize - update progress
		window.addEventListener('resize', function(event){
			if(progressEvent) return;
			progressEvent = true;
			(!window.requestAnimationFrame) ? setTimeout(function(){resetProgressIndicator();}, 250) : window.requestAnimationFrame(resetProgressIndicator);
		});

		function setProgressIndicator() {
			progressInfo['top'] = readingIndicatorContent.getBoundingClientRect().top;
			if(progressInfo['height'] <= progressInfo['window']) {
				// short content - hide progress indicator
				Util.removeClass(readingIndicator, progressInfo['class']);
				progressEvent = false;
				return;
			}
			// get new progress and update element
			Util.addClass(readingIndicator, progressInfo['class']);
			var value = (progressInfo['top'] >= 0) ? 0 : 100*(0 - progressInfo['top'])/(progressInfo['height'] - progressInfo['window']);
			readingIndicator.setAttribute('value', value);
			if(!progressIsSupported && progressFallback) progressFallback.style.width = value+'%';
			progressEvent = false;
		};

		function resetProgressIndicator() {
			progressInfo['height'] = readingIndicatorContent.offsetHeight;
			progressInfo['window'] = window.innerHeight;
			setProgressIndicator();
		};
	}
}());
// File#: _1_sticky-hero
(function() {
	var StickyBackground = function(element) {
		this.element = element;
		this.scrollingElement = this.element.getElementsByClassName('sticky-hero__content')[0];
		this.nextElement = this.element.nextElementSibling;
		this.scrollingTreshold = 0;
		this.nextTreshold = 0;
		initStickyEffect(this);
	};

	function initStickyEffect(element) {
		var observer = new IntersectionObserver(stickyCallback.bind(element), { threshold: [0, 0.1, 1] });
		observer.observe(element.scrollingElement);
		if(element.nextElement) observer.observe(element.nextElement);
	};

	function stickyCallback(entries, observer) {
		var threshold = entries[0].intersectionRatio.toFixed(1);
		(entries[0].target ==  this.scrollingElement)
			? this.scrollingTreshold = threshold
			: this.nextTreshold = threshold;

		Util.toggleClass(this.element, 'sticky-hero--media-is-fixed', (this.nextTreshold > 0 || this.scrollingTreshold > 0));
	};


	var stickyBackground = document.getElementsByClassName('js-sticky-hero'),
		intersectionObserverSupported = ('IntersectionObserver' in window && 'IntersectionObserverEntry' in window && 'intersectionRatio' in window.IntersectionObserverEntry.prototype);
	if(stickyBackground.length > 0 && intersectionObserverSupported) { // if IntersectionObserver is not supported, animations won't be triggeres
		for(var i = 0; i < stickyBackground.length; i++) {
			(function(i){ // if animations are enabled -> init the StickyBackground object
        if( Util.hasClass(stickyBackground[i], 'sticky-hero--overlay-layer') || Util.hasClass(stickyBackground[i], 'sticky-hero--scale')) new StickyBackground(stickyBackground[i]);
      })(i);
		}
	}
}());
// File#: _1_sub-navigation
// Usage: codyhouse.co/license
(function() {
    var SideNav = function(element) {
      this.element = element;
      this.control = this.element.getElementsByClassName('js-subnav__control');
      this.navList = this.element.getElementsByClassName('js-subnav__wrapper');
      this.closeBtn = this.element.getElementsByClassName('js-subnav__close-btn');
      this.firstFocusable = getFirstFocusable(this);
      this.showClass = 'subnav__wrapper--is-visible';
      this.collapsedLayoutClass = 'subnav--collapsed';
      initSideNav(this);
    };
  
    function getFirstFocusable(sidenav) { // get first focusable element inside the subnav
      if(sidenav.navList.length == 0) return;
      var focusableEle = sidenav.navList[0].querySelectorAll('[href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), button:not([disabled]), iframe, object, embed, [tabindex]:not([tabindex="-1"]), [contenteditable], audio[controls], video[controls], summary'),
          firstFocusable = false;
      for(var i = 0; i < focusableEle.length; i++) {
        if( focusableEle[i].offsetWidth || focusableEle[i].offsetHeight || focusableEle[i].getClientRects().length ) {
          firstFocusable = focusableEle[i];
          break;
        }
      }
  
      return firstFocusable;
    };
  
    function initSideNav(sidenav) {
      checkSideNavLayout(sidenav); // switch from --compressed to --expanded layout
      initSideNavToggle(sidenav); // mobile behavior + layout update on resize
    };
  
    function initSideNavToggle(sidenav) { 
      // custom event emitted when window is resized
      sidenav.element.addEventListener('update-sidenav', function(event){
        checkSideNavLayout(sidenav);
      });
  
      // mobile only
      if(sidenav.control.length == 0 || sidenav.navList.length == 0) return;
      sidenav.control[0].addEventListener('click', function(event){ // open sidenav
        openSideNav(sidenav, event);
      });
      sidenav.element.addEventListener('click', function(event) { // close sidenav when clicking on close button/bg layer
        if(event.target.closest('.js-subnav__close-btn') || Util.hasClass(event.target, 'js-subnav__wrapper')) {
          closeSideNav(sidenav, event);
        }
      });
    };
  
    function openSideNav(sidenav, event) { // open side nav - mobile only
      event.preventDefault();
      sidenav.selectedTrigger = event.target;
      event.target.setAttribute('aria-expanded', 'true');
      Util.addClass(sidenav.navList[0], sidenav.showClass);
      sidenav.navList[0].addEventListener('transitionend', function cb(event){
        sidenav.navList[0].removeEventListener('transitionend', cb);
        sidenav.firstFocusable.focus();
      });
    };
  
    function closeSideNav(sidenav, event, bool) { // close side sidenav - mobile only
      if( !Util.hasClass(sidenav.navList[0], sidenav.showClass) ) return;
      if(event) event.preventDefault();
      Util.removeClass(sidenav.navList[0], sidenav.showClass);
      if(!sidenav.selectedTrigger) return;
      sidenav.selectedTrigger.setAttribute('aria-expanded', 'false');
      if(!bool) sidenav.selectedTrigger.focus();
      sidenav.selectedTrigger = false; 
    };
  
    function checkSideNavLayout(sidenav) { // switch from --compressed to --expanded layout
      var layout = getComputedStyle(sidenav.element, ':before').getPropertyValue('content').replace(/\'|"/g, '');
      if(layout != 'expanded' && layout != 'collapsed') return;
      Util.toggleClass(sidenav.element, sidenav.collapsedLayoutClass, layout != 'expanded');
    };
    
    var sideNav = document.getElementsByClassName('js-subnav'),
      SideNavArray = [],
      j = 0;
    if( sideNav.length > 0) {
      for(var i = 0; i < sideNav.length; i++) {
        var beforeContent = getComputedStyle(sideNav[i], ':before').getPropertyValue('content');
        if(beforeContent && beforeContent !='' && beforeContent !='none') {
          j = j + 1;
        }
        (function(i){SideNavArray.push(new SideNav(sideNav[i]));})(i);
      }
  
      if(j > 0) { // on resize - update sidenav layout
        var resizingId = false,
          customEvent = new CustomEvent('update-sidenav');
        window.addEventListener('resize', function(event){
          clearTimeout(resizingId);
          resizingId = setTimeout(doneResizing, 300);
        });
  
        function doneResizing() {
          for( var i = 0; i < SideNavArray.length; i++) {
            (function(i){SideNavArray[i].element.dispatchEvent(customEvent)})(i);
          };
        };
  
        (window.requestAnimationFrame) // init table layout
          ? window.requestAnimationFrame(doneResizing)
          : doneResizing();
      }
  
      // listen for key events
      window.addEventListener('keyup', function(event){
        if( (event.keyCode && event.keyCode == 27) || (event.key && event.key.toLowerCase() == 'escape' )) {// listen for esc key - close navigation on mobile if open
          for(var i = 0; i < SideNavArray.length; i++) {
            (function(i){closeSideNav(SideNavArray[i], event);})(i);
          };
        }
        if( (event.keyCode && event.keyCode == 9) || (event.key && event.key.toLowerCase() == 'tab' )) { // listen for tab key - close navigation on mobile if open when nav loses focus
          if( document.activeElement.closest('.js-subnav__wrapper')) return;
          for(var i = 0; i < SideNavArray.length; i++) {
            (function(i){closeSideNav(SideNavArray[i], event, true);})(i);
          };
        }
      });
    }
  }());
// File#: _2_pricing-table
(function() {
	// NOTE: you need the js code only when using the --has-switch variation of the pricing table
	// default version does not require js
	var pTable = document.getElementsByClassName('js-p-table--has-switch');
	if(pTable.length > 0) {
		for(var i = 0; i < pTable.length; i++) {
			(function(i){ addPTableEvent(pTable[i]);})(i);
		}

		function addPTableEvent(element) {
			var pSwitch = element.getElementsByClassName('js-p-table__switch')[0];
			if(pSwitch) {
				pSwitch.addEventListener('change', function(event) {
          Util.toggleClass(element, 'p-table--yearly', (event.target.value == 'yearly'));
				});
			}
		}
	}
}());
// File#: _1_vertical-timeline
// Usage: codyhouse.co/license
(function() {
    var VTimeline = function(element) {
      this.element = element;
      this.sections = this.element.getElementsByClassName('js-v-timeline__section');
      this.animate = this.element.getAttribute('data-animation') && this.element.getAttribute('data-animation') == 'on' ? true : false;
      this.animationClass = 'v-timeline__section--animate';
      this.animationDelta = '-150px';
      initVTimeline(this);
    };
  
    function initVTimeline(element) {
      if(!element.animate) return;
      for(var i = 0; i < element.sections.length; i++) {
        var observer = new IntersectionObserver(vTimelineCallback.bind(element, i),
        {rootMargin: "0px 0px "+element.animationDelta+" 0px"});
        observer.observe(element.sections[i]);
      }
    };
  
    function vTimelineCallback(index, entries, observer) {
      if(entries[0].isIntersecting) {
        Util.addClass(this.sections[index], this.animationClass);
        observer.unobserve(this.sections[index]);
      } 
    };
  
    //initialize the VTimeline objects
    var timelines = document.querySelectorAll('.js-v-timeline'),
      intersectionObserverSupported = ('IntersectionObserver' in window && 'IntersectionObserverEntry' in window && 'intersectionRatio' in window.IntersectionObserverEntry.prototype),
      reducedMotion = Util.osHasReducedMotion();
    if( timelines.length > 0) {
      for( var i = 0; i < timelines.length; i++) {
        if(intersectionObserverSupported && !reducedMotion) (function(i){new VTimeline(timelines[i]);})(i);
        else timelines[i].removeAttribute('data-animation');
      }
    }
  }());