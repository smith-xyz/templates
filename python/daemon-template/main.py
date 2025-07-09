#!/usr/bin/env python3
"""
Python Daemon Template
A proper daemon implementation following Unix daemon conventions
"""

import os
import sys
import signal
import time
import atexit
import logging
from pathlib import Path
import src.utils as utils

# Configuration
DAEMON_NAME = "daemon-template"
PID_FILE = f"/tmp/{DAEMON_NAME}.pid"
LOG_FILE = f"/tmp/{DAEMON_NAME}.log"
WORKING_DIR = "/"


class Daemon:
    """
    A generic daemon class.
    
    Usage: subclass the Daemon class and override the run() method
    """
    
    def __init__(self, pidfile, stdin='/dev/null', stdout='/dev/null', stderr='/dev/null'):
        self.stdin = stdin
        self.stdout = stdout
        self.stderr = stderr
        self.pidfile = pidfile
        
    def daemonize(self):
        """
        Do the UNIX double-fork magic, see Stevens' "Advanced 
        Programming in the UNIX Environment" for details (ISBN 0201563177)
        """
        try:
            # First fork
            pid = os.fork()
            if pid > 0:
                # Exit first parent
                sys.exit(0)
        except OSError as err:
            sys.stderr.write(f"Fork #1 failed: {err}\n")
            sys.exit(1)
        
        # Decouple from parent environment
        os.chdir(WORKING_DIR)
        os.setsid()
        os.umask(0)
        
        # Second fork
        try:
            pid = os.fork()
            if pid > 0:
                # Exit from second parent
                sys.exit(0)
        except OSError as err:
            sys.stderr.write(f"Fork #2 failed: {err}\n")
            sys.exit(1)
        
        # Redirect standard file descriptors
        sys.stdout.flush()
        sys.stderr.flush()
        
        si = open(self.stdin, 'r')
        so = open(self.stdout, 'a+')
        se = open(self.stderr, 'a+')
        
        os.dup2(si.fileno(), sys.stdin.fileno())
        os.dup2(so.fileno(), sys.stdout.fileno())
        os.dup2(se.fileno(), sys.stderr.fileno())
        
        # Write pidfile
        atexit.register(self.delpid)
        
        pid = str(os.getpid())
        with open(self.pidfile, 'w+') as f:
            f.write(pid + '\n')
    
    def delpid(self):
        """Remove the pidfile"""
        if os.path.exists(self.pidfile):
            os.remove(self.pidfile)
    
    def start(self):
        """Start the daemon"""
        # Check for a pidfile to see if the daemon already runs
        try:
            with open(self.pidfile, 'r') as pf:
                pid = int(pf.read().strip())
        except (IOError, ValueError):
            pid = None
        
        if pid:
            message = f"Pidfile {self.pidfile} already exists. Daemon already running?\n"
            sys.stderr.write(message)
            sys.exit(1)
        
        # Start the daemon
        self.daemonize()
        self.run()
    
    def stop(self):
        """Stop the daemon"""
        # Get the pid from the pidfile
        try:
            with open(self.pidfile, 'r') as pf:
                pid = int(pf.read().strip())
        except (IOError, ValueError):
            pid = None
        
        if not pid:
            message = f"Pidfile {self.pidfile} does not exist. Daemon not running?\n"
            sys.stderr.write(message)
            return
        
        # Try killing the daemon process
        try:
            while True:
                os.kill(pid, signal.SIGTERM)
                time.sleep(0.1)
        except OSError as err:
            if err.errno == 3:  # No such process
                if os.path.exists(self.pidfile):
                    os.remove(self.pidfile)
            else:
                print(str(err))
                sys.exit(1)
    
    def restart(self):
        """Restart the daemon"""
        self.stop()
        self.start()
    
    def run(self):
        """
        You should override this method when you subclass Daemon.
        It will be called after the process has been daemonized by 
        start() or restart().
        """
        pass


class MyDaemon(Daemon):
    """Custom daemon implementation"""
    
    def __init__(self):
        super().__init__(PID_FILE, stdout=LOG_FILE, stderr=LOG_FILE)
        self.setup_logging()
        self.setup_signal_handlers()
        self.running = True
    
    def setup_logging(self):
        """Setup logging configuration"""
        logging.basicConfig(
            level=logging.INFO,
            format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
            handlers=[
                logging.FileHandler(LOG_FILE),
                logging.StreamHandler()
            ]
        )
        self.logger = logging.getLogger(DAEMON_NAME)
    
    def setup_signal_handlers(self):
        """Setup signal handlers for graceful shutdown"""
        def signal_handler(signum, frame):
            self.logger.info(f"Received signal {signum}, shutting down...")
            self.running = False
        
        signal.signal(signal.SIGTERM, signal_handler)
        signal.signal(signal.SIGINT, signal_handler)
    
    def run(self):
        """Main daemon loop"""
        self.logger.info(f"{DAEMON_NAME} started with PID {os.getpid()}")
        
        while self.running:
            # Main daemon work goes here
            self.logger.info("Daemon is running...")
            
            # Example: Use the utils module
            try:
                # Add your daemon logic here
                pass
            except Exception as e:
                self.logger.error(f"Error in daemon loop: {e}")
            
            # Sleep for a bit before next iteration
            time.sleep(30)
        
        self.logger.info(f"{DAEMON_NAME} stopped")


def main():
    """Main entry point"""
    daemon = MyDaemon()
    
    if len(sys.argv) == 2:
        command = sys.argv[1]
        
        if command == 'start':
            daemon.start()
        elif command == 'stop':
            daemon.stop()
        elif command == 'restart':
            daemon.restart()
        elif command == 'status':
            # Check if daemon is running
            try:
                with open(PID_FILE, 'r') as pf:
                    pid = int(pf.read().strip())
                    try:
                        os.kill(pid, 0)  # Check if process exists
                        print(f"{DAEMON_NAME} is running with PID {pid}")
                    except OSError:
                        print(f"{DAEMON_NAME} is not running")
            except (IOError, ValueError):
                print(f"{DAEMON_NAME} is not running")
        else:
            print(f"Unknown command: {command}")
            print("Usage: {0} start|stop|restart|status".format(sys.argv[0]))
            sys.exit(2)
    else:
        print("Usage: {0} start|stop|restart|status".format(sys.argv[0]))
        sys.exit(2)


if __name__ == "__main__":
    main() 