use std::time::Duration;
use signal_hook::consts::SIGTERM;
use signal_hook_tokio::Signals;
use futures::stream::StreamExt;
use tokio::time::{sleep, interval};
use tracing::{info, warn, error};

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Initialize logging
    tracing_subscriber::fmt()
        .with_env_filter("info")
        .init();

    info!("Starting daemon...");

    // Set up signal handling
    let signals = Signals::new(&[SIGTERM])?;

    // Spawn signal handling task
    let signal_task = tokio::spawn(handle_signals(signals));

    // Create shutdown channel
    let (shutdown_tx, mut shutdown_rx) = tokio::sync::mpsc::channel::<()>(1);

    // Main daemon work loop
    let daemon_task = tokio::spawn(async move {
        let mut tick_interval = interval(Duration::from_secs(10));
        let mut counter = 0;

        info!("Daemon is running...");
        
        loop {
            tokio::select! {
                _ = tick_interval.tick() => {
                    counter += 1;
                    info!("Daemon tick #{} - performing work...", counter);
                    
                    // Simulate some work
                    match perform_work(counter).await {
                        Ok(_) => info!("Work completed successfully"),
                        Err(e) => error!("Work failed: {}", e),
                    }
                }
                _ = shutdown_rx.recv() => {
                    info!("Shutdown signal received, stopping daemon...");
                    break;
                }
            }
        }
    });

    // Wait for either the signal handler or daemon task to complete
    tokio::select! {
        _ = signal_task => {
            info!("Signal received, initiating shutdown...");
            let _ = shutdown_tx.send(()).await;
        }
        _ = daemon_task => {
            info!("Daemon task completed");
        }
    }

    info!("Daemon shutdown complete");
    Ok(())
}

async fn handle_signals(mut signals: Signals) {
    while let Some(signal) = signals.next().await {
        match signal {
            SIGTERM => {
                info!("Received SIGTERM, preparing to shutdown...");
                break;
            }
            _ => {
                warn!("Received unexpected signal: {}", signal);
            }
        }
    }
}

async fn perform_work(iteration: u64) -> Result<(), Box<dyn std::error::Error>> {
    // Simulate some async work
    sleep(Duration::from_millis(100)).await;
    
    // Example: periodic maintenance, health checks, data processing, etc.
    if iteration % 5 == 0 {
        info!("Performing maintenance task at iteration {}", iteration);
    }
    
    Ok(())
}
