(print (not nil))
(print (not t))
(print (setq x nil))
(print (not x))
(print (not 'x))
(print (apply 'not '(x)))
(print (apply 'not (cons x nil)))
